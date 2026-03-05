package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID		int		`json:"id"`
	Phone	string	`json:"phone"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, phone string) (string, error) {
	expiresIn := os.Getenv("JWT_EXPIRES")

	duration, err := time.ParseDuration(expiresIn)
	if err != nil {
		duration = 24 * time.Hour
	}

	claims := Claims{
		ID: id,
		Phone: phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}