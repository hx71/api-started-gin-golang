package jwtauth

import "github.com/golang-jwt/jwt"

type Usecase interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}
