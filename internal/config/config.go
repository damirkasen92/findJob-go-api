package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDSN     string
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		DBDSN:     os.Getenv("DB_DSN"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
