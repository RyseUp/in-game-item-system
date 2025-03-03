package config

import (
	"log"

	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PostgresSQL string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PostgresSQL: os.Getenv("POSTGRES_DB"),
	}
}
