package auth

import "github.com/hx71/api-started-gin-golang/app/dto"

type Usecase interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterValidation) error
	FindByEmail(email string) bool
}
