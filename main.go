package main

import (
	"github.com/gin-gonic/gin"
	// при загрузке приложения сразу выполняется init из config.go
	"github.com/txzy2/simple-api/internal/config"
)

func main() {
	router := gin.Default()
	// Сетапим все доступные роуты
	config.SetupRoutes(router)
	router.Run(":" + config.AppConfig.Port)
}
