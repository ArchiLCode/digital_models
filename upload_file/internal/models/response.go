package models

// FileResponse представляет информацию о загруженном файле
type FileResponse struct {
	Filename     string `json:"filename"`
	Path         string `json:"path"`
	AbsolutePath string `json:"absolutePath"`
	Size         int64  `json:"size"`
	MimeType     string `json:"mimeType"`
}

// UploadResponse представляет ответ на запрос загрузки
type UploadResponse struct {
	Message string         `json:"message"`
	Files   []FileResponse `json:"files"`
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}
