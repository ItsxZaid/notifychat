package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT is not set")
	}

	return &Config{
		DatabaseURL: dbUrl,
		Port:        port,
	}, nil
}
