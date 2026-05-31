package test

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/mock"
)

type MockMenuUsecase struct {
	mock.Mock
}

func (m *MockMenuUsecase) Create(req dto.MenuCreateValidation) error {
	args := m.Called(req)
	return args.Error(0)
}
func (m *MockMenuUsecase) Show(id string) models.Menus {
	args := m.Called(id)
	return args.Get(0).(models.Menus)
}
func (m *MockMenuUsecase) Update(req dto.MenuCreateValidation) error {
	args := m.Called(req)
	return args.Error(0)
}
func (m *MockMenuUsecase) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockMenuUsecase) Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response {
	args := m.Called(ctx, pagination)
	return args.Get(0).(response.Response)
}
