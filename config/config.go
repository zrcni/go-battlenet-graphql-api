package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Setup() {
	setupEnv()
}

// setupEnv Sets environment variables from .env file
func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
