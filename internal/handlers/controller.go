package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/txzy2/simple-api/internal/status"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Controller struct{}

func (c *Controller) SuccessResponse(ctx *gin.Context, message string, data ...any) {
	response := Response{
		Success: true,
		Message: message,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	ctx.JSON(200, response)
}

func (c *Controller) ErrorResponse(ctx *gin.Context, code int, message ...string) {
	errorMessage := status.GetErrorMessage(code)

	if len(message) > 0 && message[0] != "" {
		errorMessage = message[0]
	}

	ctx.JSON(code, Response{
		Success: false,
		Message: errorMessage,
	})
}
