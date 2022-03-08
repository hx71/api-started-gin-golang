package service

import (
	"fmt"
	"srp-golang/repository"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTService is a contract of what jwtService can do
type JWTService interface {
	FindByEmails(email string) string
	ValidateToken(signedToken string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey       string
	issuer          string
	ExpirationHours int64
	userRepository  repository.UserRepository
}

type jwtCustomClaim struct {
	Email string
	jwt.StandardClaims
}

func NewJWTService(userRep repository.UserRepository) JWTService {
	return &jwtService{
		userRepository: userRep,
	}
}

func (service *jwtService) FindByEmails(email string) string {
	user := service.userRepository.FindByEmail(email)
	claims := &jwtCustomClaim{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(service.ExpirationHours)).Unix(),
			Issuer:    service.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}

	return signedToken
}

func (service *jwtService) ValidateToken(signedToken string) (*jwt.Token, error) {
	return jwt.Parse(signedToken, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

	// token, err := jwt.ParseWithClaims(
	// 	signedToken,
	// 	&jwtCustomClaim{},
	// 	func(token *jwt.Token) (interface{}, error) {
	// 		return []byte(service.secretKey), nil
	// 	},
	// )
	// if err != nil {
	// 	return token, err
	// }

	// claims, ok := token.Claims.(*jwtCustomClaim)
	// if !ok {
	// 	err = errors.New("Couldn't parse claims")
	// 	return token, err
	// }

	// if claims.ExpiresAt < time.Now().Local().Unix() {
	// 	err = errors.New("JWT is expired")
	// 	return token, err
	// }

	// return token, err

	// return jwt.Parse(token, func(claims *jwt.Token) (interface{}, error) {
	// 	token, err := jwt.ParseWithClaims(
	// 		signedToken,
	// 		&JwtClaim{},
	// 		func(token *jwt.Token) (interface{}, error) {
	// 			return []byte(service.SecretKey), nil
	// 		},
	// 	)
	// 	// 	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
	// 	// 		return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
	// 	// 	}
	// 	// 	return token, nil
	// 	// 	// return []byte(service.secretKey), nil
	// })
}
