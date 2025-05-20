package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TokenKey string
}

func MustLoad() *Config {
	// Загружаем переменные окружения из файла gateway.env
	err := godotenv.Load("/app/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	config := &Config{
		TokenKey: os.Getenv("TOKEN_KEY"),
	}

	return config
}
