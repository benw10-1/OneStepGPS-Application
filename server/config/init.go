package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
