package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{Token: os.Getenv("TOKEN")}
}
