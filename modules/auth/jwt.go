package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtAuth struct {
	secretKey []byte
}

const TIME_TOKEN_30m = 20 * time.Minute

func NewJWTAuth(secretKey string) JWTAuth {
	return &jwtAuth{secretKey: []byte(secretKey)}
}

func (j *jwtAuth) GenerateToken(userID string) (string, error) {

	expirationTime := time.Now().Add(TIME_TOKEN_30m)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *jwtAuth) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
