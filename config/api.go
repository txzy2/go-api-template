package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/txzy2/simple-api/internal/controllers/v1"
	"github.com/txzy2/simple-api/internal/middleware"
	"github.com/txzy2/simple-api/internal/services"
	db "github.com/txzy2/simple-api/pkg/database"
)

func SetupRoutes(router *gin.Engine) {
	db, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}
	servicesProvider := services.NewProvider(db)

	// Создаем контроллеры, передавая им провайдер
	testController := v1.NewTestController()
	userController := v1.NewUserController(servicesProvider)
	incidentController := v1.NewIncidentController(servicesProvider)

	api := router.Group("/api/v1")
	{
		api.GET("/", testController.Hello)
		api.POST("/test", testController.TestError)

		userGroup := api.Group("/user")
		{
			userGroup.GET("/:id", userController.GetUserById)
			userGroup.POST("/create", userController.CreateNewUser)
		}

		incidentGroup := api.Group("/incident")
		incidentGroup.Use(middleware.TokenCheck()) // <-- подключаем middleware
		{
			incidentGroup.POST("/new", incidentController.New)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})
}
