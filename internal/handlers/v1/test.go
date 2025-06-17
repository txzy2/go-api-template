package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/txzy2/simple-api/internal/handlers"
)

type TestController struct {
	*handlers.Controller
}

func NewTestController() *TestController {
	return &TestController{
		Controller: &handlers.Controller{},
	}
}

func (c *TestController) Hello(ctx *gin.Context) {
	data := map[string]any{
		"user": "John",
		"age":  30,
	}
	c.SuccessResponse(ctx, "Hello World", data)
}

func (c *TestController) TestError(ctx *gin.Context) {
	c.ErrorResponse(ctx, 400, "Invalid input")
}
