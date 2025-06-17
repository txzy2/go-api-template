package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/txzy2/simple-api/internal/controllers/v1"
	"github.com/txzy2/simple-api/internal/services"
)

func SetupRoutes(router *gin.Engine) {
	// Создаем единый провайдер сервисов
	servicesProvider := services.NewProvider()

	// Создаем контроллеры, передавая им провайдер
	testController := v1.NewTestController()
	userController := v1.NewUserController(servicesProvider)

	api := router.Group("/api")
	{
		api.GET("/", testController.Hello)
		api.POST("/test", testController.TestError)

		userGroup := api.Group("/user")
		{
			userGroup.GET("/:id", userController.GetUserById)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})
}
