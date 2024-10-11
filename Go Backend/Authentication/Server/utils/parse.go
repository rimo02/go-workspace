package utils

import (
	"auth-go/models"
	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("my-secret-key"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
