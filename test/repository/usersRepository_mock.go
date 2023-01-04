package repository

import (
	"errors"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UsersRepositoryMock) Index() *[]models.User {
	args := repository.Mock.Called()
	if args.Get(0) == nil {
		return nil
	} else {
		users := args.Get(0).([]models.User)
		return &users
	}

}

func (repository *UsersRepositoryMock) Create(req models.User) (*models.User, error) {
	args := repository.Mock.Called(req)
	if args.Get(0) == nil {
		return nil, errors.New("failed create users")
	} else {
		users := args.Get(0).(models.User)
		return &users, nil
	}
}

func (repository *UsersRepositoryMock) Show(id string) *models.User {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	} else {
		users := args.Get(0).(models.User)
		return &users
	}
}

func (repository *UsersRepositoryMock) Update(req models.User) (*models.User, error) {
	args := repository.Mock.Called(req)
	if args.Get(0) == nil {
		return nil, errors.New("failed update users")
	} else {
		users := args.Get(0).(models.User)
		return &users, nil
	}
}

func (repository *UsersRepositoryMock) Delete(req models.User) (*models.User, error) {
	args := repository.Mock.Called(req)
	if args.Get(0) == nil {
		return nil, errors.New("failed delete users")
	} else {
		users := args.Get(0).(models.User)
		return &users, nil
	}
}
