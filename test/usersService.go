package test

import (
	"errors"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	test "github.com/hasrulrhul/service-repository-pattern-gin-golang/test/repository"
)

type UsersService struct {
	// Repository repository.UsersRepository
	Repository test.UsersRepository
}

func (service UsersService) Get(id string) (*models.Users, error) {
	users := service.Repository.FindByID(id)
	if users == nil {
		return nil, errors.New("users not found")
	} else {
		return users, nil
	}
}
