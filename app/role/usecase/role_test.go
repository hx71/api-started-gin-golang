package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/role/usecase"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

// MockRepository is a mock for role.Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(model models.Role) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockRepository) Show(id string) models.Role {
	args := m.Called(id)
	return args.Get(0).(models.Role)
}

func (m *MockRepository) Update(req models.Role) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockRepository) Delete(model models.Role) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockRepository) Pagination(pagination *response.Pagination) (response.RepositoryResult, int) {
	args := m.Called(pagination)
	return args.Get(0).(response.RepositoryResult), args.Int(1)
}

func TestRoleUsecase_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewRoleUsecase(mockRepo)

	tests := []struct {
		name    string
		req     dto.RoleCreateValidation
		mockErr error
		wantErr bool
	}{
		{
			name: "Success Create Role",
			req: dto.RoleCreateValidation{
				Name: "Admin",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Failed Create Role",
			req: dto.RoleCreateValidation{
				Name: "User",
			},
			mockErr: errors.New("db error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Expected model is created with an ID inside the usecase, so we use mock.Anything
			mockRepo.On("Create", mock.AnythingOfType("models.Role")).Return(tt.mockErr).Once()

			err := uc.Create(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRoleUsecase_Show(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewRoleUsecase(mockRepo)

	expectedRole := models.Role{ID: "123", Name: "Admin"}

	tests := []struct {
		name     string
		searchID string
		mockResp models.Role
		wantResp models.Role
	}{
		{
			name:     "Success Show Role",
			searchID: "123",
			mockResp: expectedRole,
			wantResp: expectedRole,
		},
		{
			name:     "Role Not Found",
			searchID: "999",
			mockResp: models.Role{},
			wantResp: models.Role{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Show", tt.searchID).Return(tt.mockResp).Once()

			res := uc.Show(tt.searchID)
			assert.Equal(t, tt.wantResp.ID, res.ID)
			assert.Equal(t, tt.wantResp.Name, res.Name)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRoleUsecase_Update(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewRoleUsecase(mockRepo)

	tests := []struct {
		name    string
		req     dto.RoleCreateValidation
		mockErr error
		wantErr bool
	}{
		{
			name: "Success Update Role",
			req: dto.RoleCreateValidation{
				ID:   "123",
				Name: "Super Admin",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Failed Update Role",
			req: dto.RoleCreateValidation{
				ID:   "123",
				Name: "Error Admin",
			},
			mockErr: errors.New("update error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Update", mock.AnythingOfType("models.Role")).Return(tt.mockErr).Once()

			err := uc.Update(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRoleUsecase_Delete(t *testing.T) {
	mockRepo := new(MockRepository)
	uc := usecase.NewRoleUsecase(mockRepo)

	tests := []struct {
		name    string
		model   models.Role
		mockErr error
		wantErr bool
	}{
		{
			name: "Success Delete Role",
			model: models.Role{
				ID: "123",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "Failed Delete Role",
			model: models.Role{
				ID: "999",
			},
			mockErr: errors.New("delete error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("Delete", tt.model).Return(tt.mockErr).Once()

			err := uc.Delete(tt.model)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
