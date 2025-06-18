package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(consumerID int64, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"consumer_id": consumerID,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ParseJWT(tokenStr string, secretKey string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid claims")
	}

	idFloat, ok := claims["consumer_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid consumer_id in token")
	}

	return int64(idFloat), nil
}
