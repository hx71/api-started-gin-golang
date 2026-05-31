package test

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/mock"
)

type MockAuditLogUsecase struct {
	mock.Mock
}

func (m *MockAuditLogUsecase) Create(req dto.AuditLogCreateValidation) error {
	args := m.Called(req)
	return args.Error(0)
}
func (m *MockAuditLogUsecase) Show(id string) models.AuditLog {
	args := m.Called(id)
	return args.Get(0).(models.AuditLog)
}
func (m *MockAuditLogUsecase) Delete(model models.AuditLog) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockAuditLogUsecase) Pagination(ctx *gin.Context, pagination *response.Pagination) response.Response {
	args := m.Called(ctx, pagination)
	return args.Get(0).(response.Response)
}
