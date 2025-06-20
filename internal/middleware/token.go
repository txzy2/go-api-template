package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	timestampTolerance = 400
)

type HeaderAuthData struct {
	XTimestamp string `header:"x-timestamp" binding:"required,numeric"`
	XSignature string `header:"x-signature" binding:"required,hexadecimal"`
}

type Request struct {
	body      string
	path      string
	method    string
	timestamp string
}

func GenerateHMAC(params *Request, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	data := params.body + params.path + params.method + params.timestamp
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headers HeaderAuthData

		if err := c.ShouldBindHeader(&headers); err != nil {
			resp := gin.H{"error": "invalid headers"}
			requiredHeaders := []string{"x-timestamp", "x-signature"}
			for _, h := range requiredHeaders {
				if c.GetHeader(h) == "" {
					resp[h] = "required"
				}
			}
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}

		// Проверка timestamp (число + не устаревший)
		ts, err := strconv.ParseInt(headers.XTimestamp, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "invalid timestamp",
				"details": "x-timestamp must be a valid Unix timestamp",
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > ts+timestampTolerance {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "expired timestamp",
				"details": "the request timestamp is too old",
			})
			c.Abort()
			return
		}

		// Проверка наличия TOKEN_SALT
		key := os.Getenv("TOKEN_SALT")
		if key == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "server misconfiguration",
				"details": "internal service error",
			})
			c.Abort()
			return
		}

		// Чтение тела запроса
		bodyBytes, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "failed to read request body",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		// Генерация и проверка HMAC
		params := Request{
			body:      string(bodyBytes),
			path:      c.Request.URL.Path,
			method:    c.Request.Method,
			timestamp: headers.XTimestamp,
		}

		expectedSignature := GenerateHMAC(&params, key)
		fmt.Printf("[DEBUG] HMAC: %s | x-signature: %+v\n", expectedSignature, headers.XSignature)
		if expectedSignature != headers.XSignature {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "invalid signature",
				"details": "verification failed",
			})
			c.Abort()
			return
		}

		// Успешная аутентификация
		c.Set("x-timestamp", headers.XTimestamp)
		c.Set("x-signature", headers.XSignature)
		c.Next()
	}
}
