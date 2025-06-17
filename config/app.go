package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/txzy2/simple-api/pkg/common"
	"github.com/txzy2/simple-api/pkg/logger"
)

type Config struct {
	// Порт api
	Port string

	// Конфиг для БД
	DB DBConfig
}

var AppConfig Config

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		Port: getEnv("PORT", "8080"),
		DB:   LoadDBConfig(),
	}

	common.SetMode(getEnv("GIN_MODE", "dev"))
}

func init() {
	loadConfig()
	logger.Init()
}
