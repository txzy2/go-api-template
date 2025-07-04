package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txzy2/simple-api/internal/controllers/status"
	"github.com/txzy2/simple-api/pkg/common"
)

type (
	Success struct {
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
		Mode    string `json:"mode"`
	}

	Fail struct {
		Error string `json:"error"`
		Mode  string `json:"mode,omitempty"`
	}
)

type Controller struct{}

func (c *Controller) SuccessResponse(ctx *gin.Context, message string, data any) {
	response := Success{
		Message: message,
		Data:    data,
		Mode:    common.Mode,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) ErrorResponse(ctx *gin.Context, code int, errMessage string) {
	errorMessage := status.GetErrorMessage(code)

	if errMessage != "" {
		errorMessage = errMessage
	}

	response := Fail{
		Error: errorMessage,
		Mode:  common.Mode,
	}

	ctx.JSON(code, response)
}
