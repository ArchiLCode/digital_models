package models

type Session struct {
}

// SessionRequest represents the request payload for user session.
// swagger:model
type SessionRequest struct {
	// Access token to invalidate
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9...
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IlJXTiJ9..."`
}

// SessionResponse represents the response payload after successful session.
// swagger:model
type SessionResponse struct {
	// ID of the newly created user
	// example: 123
	UserID string `json:"user_id" example:"123"`

	// Last name of the user
	// example: Doe
	LastName string `json:"last_name" example:"Doe"`

	// Name of the user
	// example: Joe
	Name string `json:"name" example:"Joe"`

	// Goal of the user
	// example: model
	Goal string `json:"goal" example:"model"`

	// IsAdmin
	// example: false
	IsAdmin bool `json:"is_admin" example:"false"`
}
