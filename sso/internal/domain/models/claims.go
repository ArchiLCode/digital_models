package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	JTI   string    `json:"jti"`
	ID    uuid.UUID `json:"uid"`
	Email string    `json:"email"`
	AppID int       `json:"app_id"`
	Time  int64     `json:"exp"`
	jwt.RegisteredClaims
}
