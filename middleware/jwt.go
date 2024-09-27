package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaim struct {
	Email string
	jwt.StandardClaims
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "secret"
	}
	return secretKey
}

func GenerateToken(Email string) string {
	claims := &JwtCustomClaim{
		Email,
		jwt.StandardClaims{
			// ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(60)).Unix(),
			Issuer:    getSecretKey(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(getSecretKey()), nil
	})
}
