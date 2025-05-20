package models

// RefreshTokenRequest represents the request payload for refreshing the access token.
// swagger:model
type RefreshTokenRequest struct {
	// Refresh token for obtaining a new access token
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9...
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9..."`
}

// RefreshTokenResponse represents the response payload after successfully refreshing the access token.
// swagger:model
type RefreshTokenResponse struct {
	// New access token for the user
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`

	// New refresh token for obtaining another access token
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9...
	NewRefreshToken string `json:"new_refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9..."`
}
