package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("supersecret") // use env in prod

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
