package repository

import (
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UsersRepositoryMock) FindByID(id string) *models.Users {
	args := repository.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	} else {
		users := args.Get(0).(models.Users)
		return &users
	}

}
