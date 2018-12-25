package api

import (
	"log"

	"github.com/joho/godotenv"
)

// SetupEnv Sets environment variables from .env file
func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
