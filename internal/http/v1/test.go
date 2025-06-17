package v1

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/txzy2/simple-api/internal/http"
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
	c.SuccessResponse(ctx, "Hello World", nil)
}

func (c *TestController) TestError(ctx *gin.Context) {
	c.ErrorResponse(ctx, 400, "Invalid input")
}
