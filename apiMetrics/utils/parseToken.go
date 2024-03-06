package utils

import (
	"apiMetrics/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("rwBKDrE81h6jFGdxyJ85TGwBT3i-eGIRLcsIYScT680="), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid or malformed token")
}
