package models

// LoginRequest represents the request payload for user login.
// swagger:model
type LoginRequest struct {
	// Email of the user
	// example: john.doe@example.com
	Email string `json:"email" example:"john.doe@example.com"`

	// Password of the user
	// example: secure_password
	Password string `json:"password" example:"secure_password"`

	// AppID id application
	// example: 1
	AppID int `json:"app_id" example:"1"`
}

// LoginResponse represents the response payload after successful login.
// swagger:model
type LoginResponse struct {
	// Access token for the user
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`

	// Refresh token for obtaining a new access token
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9...
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9..."`
}
