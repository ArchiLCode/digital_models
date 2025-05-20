package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config содержит настройки приложения
type Config struct {
	UploadFolder     string
	MaxFileSize      int64
	ServerPort       string
	AllowedMimeTypes []string
}

// LoadConfig загружает настройки из окружения или использует значения по умолчанию
func LoadConfig() *Config {
	// Настройки по умолчанию
	cfg := &Config{
		UploadFolder:     getEnv("UPLOAD_FOLDER", "../ui/public/uploads"),
		MaxFileSize:      getEnvAsInt64("MAX_FILE_SIZE", 16<<20), // 16 MB
		ServerPort:       getEnv("SERVER_PORT", ":8080"),
		AllowedMimeTypes: getEnvAsSlice("ALLOWED_MIME_TYPES", "image/jpeg,image/png,image/gif,image/bmp,image/webp"),
	}
	fmt.Println(cfg.UploadFolder)
	// Создаем директорию для загрузок, если не существует
	if err := os.MkdirAll(cfg.UploadFolder, 0755); err != nil {
		panic("Не удалось создать директорию для загрузок: " + err.Error())
	}

	return cfg
}

// Вспомогательные функции для получения переменных окружения
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			return value
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue string) []string {
	if valueStr, exists := os.LookupEnv(key); exists {
		return strings.Split(valueStr, ",")
	}
	return strings.Split(defaultValue, ",")
}
