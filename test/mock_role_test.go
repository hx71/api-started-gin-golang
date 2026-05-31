package test

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/mock"
)

type MockRoleUsecase struct {
	mock.Mock
}

func (m *MockRoleUsecase) Create(req dto.RoleCreateValidation) error {
	args := m.Called(req)
	return args.Error(0)
}
func (m *MockRoleUsecase) Show(id string) models.Role {
	args := m.Called(id)
	return args.Get(0).(models.Role)
}
func (m *MockRoleUsecase) Update(req dto.RoleCreateValidation) error {
	args := m.Called(req)
	return args.Error(0)
}
func (m *MockRoleUsecase) Delete(model models.Role) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockRoleUsecase) Pagination(ctx *gin.Context, pagination *response.Pagination) response.Response {
	args := m.Called(ctx, pagination)
	return args.Get(0).(response.Response)
}
