package usecase_test

import (
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/user/usecase"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct{ mock.Mock }

func (m *MockRepository) Create(model models.User) error           { return m.Called(model).Error(0) }
func (m *MockRepository) Show(id string) models.User               { return m.Called(id).Get(0).(models.User) }
func (m *MockRepository) Update(model models.User) error           { return m.Called(model).Error(0) }
func (m *MockRepository) Delete(model models.User) error           { return m.Called(model).Error(0) }
func (m *MockRepository) FindByEmail(email string) bool            { return m.Called(email).Bool(0) }
func (m *MockRepository) VerifyCredential(e, p string) interface{} { return m.Called(e, p).Get(0) }
func (m *MockRepository) Pagination(p *helpers.Pagination) (response.RepositoryResult, int) {
	args := m.Called(p)
	return args.Get(0).(response.RepositoryResult), args.Int(1)
}
func TestUserUsecase_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewUserUsecase(mockRepo)
	mockRepo.On("FindByEmail", mock.Anything).Return(false)
	mockRepo.On("Create", mock.Anything).Return(nil)
	err := uc.Create(dto.UserCreateValidation{Email: "test@example.com", Password: "123", })
	assert.NoError(t, err)
}
