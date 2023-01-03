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

func (service UsersService) Index() (*[]models.User, error) {
	users := service.Repository.Index()
	if users == nil {
		return nil, errors.New("users not found")
	} else {
		return users, nil
	}
}

func (service UsersService) Create(req models.User) (*models.User, error) {
	user, err := service.Repository.Create(req)
	if err != nil {
		return nil, errors.New("failed users create")
	} else {
		return user, nil
	}
}

func (service UsersService) Show(id string) (*models.User, error) {
	users := service.Repository.Show(id)
	if users == nil {
		return nil, errors.New("users not found")
	} else {
		return users, nil
	}
}
