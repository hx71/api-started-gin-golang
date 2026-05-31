package usecase_test

import (
	"github.com/hx71/api-started-gin-golang/app/auditlog/usecase"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct{ mock.Mock }

func (m *MockRepository) Create(model models.AuditLog) error { return m.Called(model).Error(0) }
func (m *MockRepository) Show(id string) models.AuditLog {
	return m.Called(id).Get(0).(models.AuditLog)
}
func (m *MockRepository) Delete(model models.AuditLog) error { return m.Called(model).Error(0) }
func (m *MockRepository) Pagination(p *response.Pagination) (response.RepositoryResult, int) {
	args := m.Called(p)
	return args.Get(0).(response.RepositoryResult), args.Int(1)
}
func TestAuditLogUsecase_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewAuditLogUsecase(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(nil)
	err := uc.Create(dto.AuditLogCreateValidation{ServiceName: "Test"})
	assert.NoError(t, err)
}
