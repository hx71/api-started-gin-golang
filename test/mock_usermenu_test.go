package test

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/mock"
)

type MockUserMenuUsecase struct {
	mock.Mock
}

func (m *MockUserMenuUsecase) Create(req []dto.UserMenuCreateValidation) error {
	args := m.Called(req)
	return args.Error(0)
}
func (m *MockUserMenuUsecase) Show(id string) models.UserMenus {
	args := m.Called(id)
	return args.Get(0).(models.UserMenus)
}
func (m *MockUserMenuUsecase) Update(id string, req dto.UserMenuCreateValidation) error {
	args := m.Called(id, req)
	return args.Error(0)
}
func (m *MockUserMenuUsecase) Delete(model models.UserMenus) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockUserMenuUsecase) Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response {
	args := m.Called(ctx, pagination)
	return args.Get(0).(response.Response)
}
