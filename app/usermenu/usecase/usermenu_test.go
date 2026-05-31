package usecase_test
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/hx71/api-started-gin-golang/app/usermenu/usecase"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/response"
)
type MockRepository struct { mock.Mock }
func (m *MockRepository) Create(model models.UserMenu) error { return m.Called(model).Error(0) }
func (m *MockRepository) Show(id string) models.UserMenus { return m.Called(id).Get(0).(models.UserMenus) }
func (m *MockRepository) Update(model models.UserMenu) error { return m.Called(model).Error(0) }
func (m *MockRepository) Delete(model models.UserMenus) error { return m.Called(model).Error(0) }
func (m *MockRepository) Pagination(p *helpers.Pagination) (response.RepositoryResult, int) {
	args := m.Called(p)
	return args.Get(0).(response.RepositoryResult), args.Int(1)
}
func TestUserMenuUsecase_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewUserMenuUsecase(mockRepo)
	mockRepo.On("Create", mock.Anything).Return(nil)
	err := uc.Create([]dto.UserMenuCreateValidation{{MenuID: "test"}})
	assert.NoError(t, err)
}
