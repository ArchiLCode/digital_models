package models

// RegisterRequest represents the request payload for user registration.
// swagger:model
type RegisterRequest struct {
	// Email of the user
	// example: john.doe@example.com
	Email string `json:"email" example:"john.doe@example.com"`

	// Goal of the user
	// example: model
	Goal string `json:"goal" example:"model"`

	// Last name of the user
	// example: Doe
	LastName string `json:"last_name" example:"Doe"`

	// Name of the user
	// example: Joe
	Name string `json:"name" example:"Joe"`

	// Password of the user
	// example: secure_password
	Password string `json:"password" example:"secure_password"`
}

// RegisterResponse represents the response payload after successful registration.
// swagger:model
type RegisterResponse struct {
	// ID of the newly created user
	// example: 123
	UserID string `json:"user_id" example:"123"`
}

// ErrorResponse represents an error response.
// swagger:model
type ErrorResponse struct {
	// Error message
	// example: Invalid input
	Error string `json:"error"`
}
