package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"upload-services/internal/config"
	"upload-services/internal/models"

	"github.com/google/uuid"
)

// FileService предоставляет методы для работы с файлами
type FileService struct {
	config *config.Config
}

// NewFileService создает новый экземпляр FileService
func NewFileService(config *config.Config) *FileService {
	return &FileService{
		config: config,
	}
}

// SaveFile сохраняет один файл на диск
func (s *FileService) SaveFile(file *multipart.FileHeader) (*models.FileResponse, error) {
	// Проверка MIME-типа файла
	if !s.isAllowedMimeType(file.Header.Get("Content-Type")) {
		return nil, fmt.Errorf("недопустимый тип файла: %s", file.Header.Get("Content-Type"))
	}

	// Открытие файла
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Генерация имени файла
	filename := s.generateUniqueFilename(file.Filename)
	filePath := filepath.Join(s.config.UploadFolder, filename)

	// Создание файла
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Копирование содержимого
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	// Получение абсолютного пути
	absolutePath, _ := filepath.Abs(filePath)

	// Формирование ответа
	fileURL := fmt.Sprintf("/uploads/%s", filename)

	return &models.FileResponse{
		Filename:     filename,
		Path:         fileURL,
		AbsolutePath: absolutePath,
		Size:         file.Size,
		MimeType:     file.Header.Get("Content-Type"),
	}, nil
}

// SaveMultipleFiles сохраняет несколько файлов на диск
func (s *FileService) SaveMultipleFiles(files []*multipart.FileHeader) ([]models.FileResponse, []error) {
	responses := make([]models.FileResponse, 0, len(files))
	errors := make([]error, 0)

	for _, file := range files {
		response, err := s.SaveFile(file)
		if err != nil {
			errors = append(errors, fmt.Errorf("ошибка при сохранении файла %s: %v", file.Filename, err))
		} else {
			responses = append(responses, *response)
		}
	}

	return responses, errors
}

// Проверка допустимости MIME-типа файла
func (s *FileService) isAllowedMimeType(mimeType string) bool {
	for _, allowedType := range s.config.AllowedMimeTypes {
		if mimeType == allowedType {
			return true
		}
	}
	return false
}

// Генерация уникального имени файла
func (s *FileService) generateUniqueFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	timestamp := time.Now().Format("20060102_150405")
	uniqueID := uuid.New().String()[:8]

	return fmt.Sprintf("%s_%s%s", timestamp, uniqueID, extension)
}
