package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort       string `env:"SERVER_PORT"`
	DBHost           string `env:"DB_HOST"`
	DBName           string `env:"DB_NAME"`
	DBUser           string `env:"DB_USER"`
	DBPassword       string `env:"DB_PASSWORD"`
	DBSSLMode        string `env:"DB_SSL_MODE"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresDB       string `env:"POSTGRES_DB"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatal("Unable to parse .env file")
	}

	return config
}
