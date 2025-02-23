package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBConnection string
	JWTSecret    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		DBConnection: os.Getenv("DB_CONNECTION"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}, nil
}
