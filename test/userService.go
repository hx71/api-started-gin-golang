package test

import (
	"errors"

	"github.com/hx71/api-started-gin-golang/models"
	test "github.com/hx71/api-started-gin-golang/test/repository"
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
	users, err := service.Repository.Create(req)
	if err != nil {
		return nil, err
	} else {
		return users, nil
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

func (service UsersService) Update(req models.User) (*models.User, error) {
	users, err := service.Repository.Update(req)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (service UsersService) Delete(req models.User) (*models.User, error) {
	users, err := service.Repository.Delete(req)
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}
