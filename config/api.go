package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/txzy2/simple-api/internal/handlers/v1"
)

func SetupRoutes(router *gin.Engine) {
	testController := v1.NewTestController()

	api := router.Group("/api")
	{
		api.GET("/", testController.Hello)
		api.POST("/test", testController.TestError)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})
}
