package auth

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(bankID string) (string, error) {
	claims := jwt.MapClaims{
		"bank_id": bankID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}