package models

// LogoutRequest represents the request payload for logging out the user.
// swagger:model
type LogoutRequest struct {
	// Refresh token to invalidate
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9...
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9..."`
}
