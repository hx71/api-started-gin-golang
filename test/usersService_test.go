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

func TestUsersService_Index(t *testing.T) {
	var listUsers []models.User
	var users = models.User{
		ID:   "1",
		Name: "hasrul",
	}
	listUsers = append(listUsers, users)

	usersRepository.Mock.On("Index").Return(listUsers)
	user, err := usersService.Index()
	assert.Nil(t, err)
	assert.NotNil(t, user)

}
func TestUsersService_Create(t *testing.T) {
	users := models.User{
		ID:       "1",
		Name:     "hasrul",
		Username: "hasrul",
		Password: "hasrul@123",
		Email:    "hasrul@mail.com",
	}

	usersRepository.Mock.On("Create", users).Return(users, nil)
	user, err := usersService.Create(users)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.ID, "1")
}

func TestUsersService_GetFound(t *testing.T) {
	users := models.User{
		ID:   "1",
		Name: "hasrul",
	}
	usersRepository.Mock.On("Show", "1").Return(users)

	result, err := usersService.Show("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.ID, "1")
	assert.Equal(t, result.Name, "hasrul")
}

func TestUsersService_GetNotFound(t *testing.T) {
	usersRepository.Mock.On("Show", "2").Return(nil)

	users, err := usersService.Show("2")
	assert.Nil(t, users)
	assert.NotNil(t, err)

}
