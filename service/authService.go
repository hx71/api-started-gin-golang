package service

import (
	"log"
	"srp-golang/app/models"
	"srp-golang/app/request"
	"srp-golang/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	// VerifyCredential(email string, password string) interface{}
	CreateUser(user request.RegisterValidation) models.User
	// FindByEmail(email string) models.User
	// IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) CreateUser(user request.RegisterValidation) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}
