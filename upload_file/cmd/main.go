package main

import (
	"log"

	"upload-services/internal/config"
	"upload-services/internal/handlers"
	"upload-services/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация сервисов
	fileService := services.NewFileService(cfg)

	// Инициализация роутера Gin
	router := gin.Default()
	router.MaxMultipartMemory = cfg.MaxFileSize

	// Настройка маршрутов
	uploadHandler := handlers.NewUploadHandler(fileService)
	handlers.SetupRoutes(router, uploadHandler)

	// Запуск сервера
	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)
	if err := router.Run(cfg.ServerPort); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
