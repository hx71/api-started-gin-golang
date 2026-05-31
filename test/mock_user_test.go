package test

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Create(model dto.UserCreateValidation) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockUserUsecase) Show(id string) models.User {
	args := m.Called(id)
	return args.Get(0).(models.User)
}
func (m *MockUserUsecase) Update(model dto.UserUpdateValidation) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockUserUsecase) Delete(model models.User) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockUserUsecase) FindByEmail(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}
func (m *MockUserUsecase) Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response {
	args := m.Called(ctx, pagination)
	return args.Get(0).(response.Response)
}
