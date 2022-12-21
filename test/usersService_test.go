package test

import (
	"testing"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// var usersRepository = &repository.UsersRepositoryMock{Mock: mock.Mock{}}
var usersRepository = &repository.UsersRepositoryMock{Mock: mock.Mock{}}
var usersService = UsersService{Repository: usersRepository}

func TestUsersService_GetNotFound(t *testing.T) {

	usersRepository.Mock.On("FindByID", "1").Return(nil)

	users, err := usersService.Get("1")
	assert.Nil(t, users)
	assert.NotNil(t, err)

}

func TestUsersService_GetFound(t *testing.T) {

	users := models.Users{
		ID:   "2",
		Name: "hasrul",
	}

	usersRepository.Mock.On("FindByID", "2").Return(users)
	result, err := usersService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.ID, "2")
	assert.Equal(t, result.Name, "hasrul")
}

// func TestUsersService_Show(t *testing.T) {

// 	users := models.Users{
// 		ID:   "3",
// 		Name: "rhul",
// 	}

// 	usersRepository.Mock.On("Show", "3").Return(users)

// 	result, err := usersService.Show("3")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, result.ID, "3")
// 	assert.Equal(t, result.Name, "rhul")
// }
