package service

import (
	"errors"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/repository"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
)

type UsersService struct {
	Repository repository.UsersRepository
}

func (service UsersService) Get(id string) (*models.Users, error) {
	users := service.Repository.FindByID(id)
	if users == nil {
		return nil, errors.New("users not found")
	} else {
		return users, nil
	}
}
