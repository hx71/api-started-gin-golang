package usecase_test

import (
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/menu/usecase"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct{ mock.Mock }

func (m *MockRepository) Create(model models.Menu) error { return m.Called(model).Error(0) }
func (m *MockRepository) Show(id string) models.Menus    { return m.Called(id).Get(0).(models.Menus) }
func (m *MockRepository) Update(model models.Menu) error { return m.Called(model).Error(0) }
func (m *MockRepository) Delete(id string) error         { return m.Called(id).Error(0) }
func (m *MockRepository) Pagination(p *helpers.Pagination) (response.RepositoryResult, int) {
	args := m.Called(p)
	return args.Get(0).(response.RepositoryResult), args.Int(1)
}
func TestMenuUsecase_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewMenuUsecase(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(nil)
	err := uc.Create(dto.MenuCreateValidation{Name: "Test"})
	assert.NoError(t, err)
}
