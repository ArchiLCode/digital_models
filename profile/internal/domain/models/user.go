package models

import (
	"github.com/google/uuid"
	"time"
)

// Структура User
type User struct {
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	Sex         string    `json:"sex"`
	DateOfBirth time.Time `json:"date_of_birth"`
	City        string    `json:"city"`
	Height      int       `json:"height"`
	Chest_size  int       `json:"chest_size"`
	Waist_size  int       `json:"waist_size"`
	Hip_size    int       `json:"hip_size"`
	Weight      int       `json:"weight"`
	Photos      []string  `json:"photos"`     // Список всех фотографий с URL-адресами
	AvatarURL   string    `json:"avatar_url"` // URL аватара
	UserID      uuid.UUID `json:"user_id"`
}
