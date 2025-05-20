package handlers

import (
	"github.com/gin-contrib/cors"
	"log"
	"net/http"
	"strings"

	"upload-services/internal/models"
	"upload-services/internal/services"

	"github.com/gin-gonic/gin"
)

// UploadHandler обрабатывает запросы на загрузку файлов
type UploadHandler struct {
	fileService *services.FileService
}

// NewUploadHandler создает новый экземпляр UploadHandler
func NewUploadHandler(fileService *services.FileService) *UploadHandler {
	return &UploadHandler{
		fileService: fileService,
	}
}

// UploadSingle обрабатывает загрузку одного файла
func (h *UploadHandler) UploadSingle(c *gin.Context) {
	// Получение файла из запроса
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Файл не найден в запросе"})
		return
	}

	// Сохранение файла
	fileResponse, err := h.fileService.SaveFile(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		log.Printf("Ошибка при сохранении файла: %v", err)
		return
	}

	// Формирование ответа
	response := models.UploadResponse{
		Message: "Файл успешно загружен",
		Files:   []models.FileResponse{*fileResponse},
	}

	c.JSON(http.StatusOK, response)
}

// UploadMultiple обрабатывает загрузку нескольких файлов
func (h *UploadHandler) UploadMultiple(c *gin.Context) {
	// Получение файлов из запроса
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Ошибка получения формы"})
		return
	}

	files := form.File["photos"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Файлы не найдены в запросе. Используйте поле 'photos' для загрузки."})
		return
	}

	// Сохранение файлов
	fileResponses, errors := h.fileService.SaveMultipleFiles(files)

	// Проверка наличия ошибок
	if len(errors) > 0 {
		errorMessages := make([]string, len(errors))
		for i, err := range errors {
			errorMessages[i] = err.Error()
		}

		if len(fileResponses) == 0 {
			// Если ни один файл не был загружен успешно
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: strings.Join(errorMessages, "; ")})
			return
		}

		// Логируем ошибки, но продолжаем, если есть успешно загруженные файлы
		log.Printf("Ошибки при загрузке файлов: %v", strings.Join(errorMessages, "; "))
	}

	// Формирование ответа
	response := models.UploadResponse{
		Message: "Файлы успешно загружены",
		Files:   fileResponses,
	}

	c.JSON(http.StatusOK, response)
}

// SetupRoutes настраивает маршруты и middleware
func SetupRoutes(router *gin.Engine, uploadHandler *UploadHandler) {
	// Настройка CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(corsConfig))

	// Маршруты
	router.POST("/upload", uploadHandler.UploadSingle)
	router.POST("/upload/multiple", uploadHandler.UploadMultiple)
	router.Static("/uploads", "./uploads")
}
