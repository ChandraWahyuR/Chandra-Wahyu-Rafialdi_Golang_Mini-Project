package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	UserId    uuid.UUID `json:"user_id"`
	UserEmail string    `json:"user_email"`
	Roles     string    `json:"role"`
	jwt.RegisteredClaims
}
