package models

import "google.golang.org/protobuf/types/known/timestamppb"

// GetProfileByIDResponse представляет ответ с данными профиля пользователя
type GetProfileByIDResponse struct {
	Name        string                 `json:"name,omitempty"`
	LastName    string                 `json:"last_name,omitempty"`
	Sex         string                 `json:"sex,omitempty"`
	DateOfBirth *timestamppb.Timestamp `json:"date_of_birth,omitempty"`
	City        string                 `json:"city,omitempty"`
	Height      int32                  `json:"height,omitempty"`
	ChestSize   int32                  `json:"chest_size,omitempty"`
	WaistSize   int32                  `json:"waist_size,omitempty"`
	HipSize     int32                  `json:"hip_size,omitempty"`
	Weight      int32                  `json:"weight,omitempty"`
	Photos      []string               `json:"photos,omitempty"`
	AvatarURL   string                 `json:"avatar_url,omitempty"`
}

// GetCatalogResponse представдяет ответ с карточками моделей
type GetCatalogResponse struct {
	Items      []CatalogItem `json:"items,omitempty"`
	TotalPages int32         `json:"total_pages,omitempty"`
}
