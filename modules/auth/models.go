package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type Claims struct {
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"exp"`
	jwt.StandardClaims
}
