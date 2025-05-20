package models

import "time"

// UpdateProfileRequest represents the request payload for user login.
// swagger:model
type UpdateProfileRequest struct {
	// ID of the newly created user
	// example: 123
	UserID string `json:"user_id" example:"123"`

	// The first name of the user
	// example: John
	Name string `json:"name" example:"John"`

	// The last name of the user
	// example: Doe
	LastName string `json:"last_name" example:"Doe"`

	// The gender of the user
	// example: male
	Sex string `json:"sex" example:"male"`

	// The date of birth of the user
	// example: 1990-01-01T00:00:00Z
	DateOfBirth time.Time `json:"date_of_birth" example:"1990-01-01T00:00:00Z"`

	// The city where the user lives
	// example: New York
	City string `json:"city" example:"New York"`

	// The height of the user in centimeters
	// example: 180
	Height int32 `json:"height" example:"180"`

	// The chest size of the user in centimeters
	// example: 95
	ChestSize int32 `json:"chest_size" example:"95"`

	// The waist size of the user in centimeters
	// example: 80
	WaistSize int32 `json:"waist_size" example:"80"`

	// The hip size of the user in centimeters
	// example: 90
	HipSize int32 `json:"hip_size" example:"90"`

	// The weight of the user in kilograms
	// example: 75
	Weight int32 `json:"weight" example:"75"`

	// List of URLs to the user's photos
	// example: ["https://example.com/photo1.jpg", "https://example.com/photo2.jpg"]
	Photos []string `json:"photos" example:"https://example.com/photo1.jpg,https://example.com/photo2.jpg"`

	// URL to the user's avatar
	// example: https://example.com/avatar.jpg
	AvatarUrl string `json:"avatar_url" example:"https://example.com/avatar.jpg"`
}
