package test

import (
	"github.com/golang-jwt/jwt"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/stretchr/testify/mock"
)

type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) VerifyCredential(email string, password string) interface{} {
	args := m.Called(email, password)
	return args.Get(0)
}

func (m *MockAuthUsecase) CreateUser(user dto.RegisterValidation) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAuthUsecase) FindByEmail(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

type MockJWTUsecase struct {
	mock.Mock
}

func (m *MockJWTUsecase) GenerateToken(userID string) string {
	args := m.Called(userID)
	return args.String(0)
}

func (m *MockJWTUsecase) ValidateToken(token string) (*jwt.Token, error) {
	args := m.Called(token)
	if args.Get(0) != nil {
		return args.Get(0).(*jwt.Token), args.Error(1)
	}
	return nil, args.Error(1)
}
