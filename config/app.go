package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/txzy2/simple-api/pkg/common"
)

type Config struct {
	// Порт api
	Port string
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
		log.Println("Warning: .env file not found, using environment variables")
	}

	AppConfig = Config{
		Port: getEnv("PORT", "8080"),
	}

	common.SetMode(getEnv("GIN_MODE", "dev"))
}

func init() {
	loadConfig()
}
