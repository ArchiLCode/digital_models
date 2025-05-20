package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateProfileRequest struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	UserID   string `json:"user_id,omitempty"`
}

// GetProfileByIDRequest представляет запрос на получение профиля по ID
type GetProfileByIDRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

// UpdateProfileRequest представляет запрос на обновление профиля
type UpdateProfileRequest struct {
	Name        string    `json:"name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Sex         string    `json:"sex,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	City        string    `json:"city"`
	Height      int32     `json:"height"`
	ChestSize   int32     `json:"chest_size"`
	WaistSize   int32     `json:"waist_size"`
	HipSize     int32     `json:"hip_size"`
	Weight      int32     `json:"weight"`
	Photos      []string  `json:"photos"`
	AvatarURL   string    `json:"avatar_url"`
	UserID      string    `json:"user_id,omitempty"`
}

// DeleteProfileRequest представляет запрос на удаление профиля
type DeleteProfileRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetCatalogRequest struct {
	DateFrom   time.Time
	DateTo     time.Time
	Sex        string
	HeightFrom int32
	HeightTo   int32
	WeightFrom int32
	WeightTo   int32
	Page       int32
	Limit      int32
}
