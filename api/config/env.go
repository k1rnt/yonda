package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvFile() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
