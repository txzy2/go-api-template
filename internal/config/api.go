package config

import (
	"github.com/gin-gonic/gin"

	"github.com/txzy2/simple-api/internal/handlers"
	v1 "github.com/txzy2/simple-api/internal/handlers/v1"
)

type BaseController struct {
	*handlers.Controller
}

func SetupRoutes(router *gin.Engine) {
	testController := v1.NewTestController()
	baseController := &BaseController{}

	router.NoRoute(func(c *gin.Context) {
		baseController.ErrorResponse(c, 404, "")
	})

	api := router.Group("/api")
	{
		api.GET("/", testController.Hello)
		api.POST("/test", testController.TestError)
	}
}
