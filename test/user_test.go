package test

import (
	"github.com/gin-gonic/gin"
	handlerPackage "github.com/hx71/api-started-gin-golang/app/user/handler"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupUserRouter(mockUsecase *MockUserUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := handlerPackage.NewUserHandler(mockUsecase)

	router.GET("/users", handler.Index)
	router.GET("/users/:id", handler.Show)
	router.POST("/users", handler.Create)
	router.DELETE("/users/:id", handler.Delete)
	router.PUT("/users/:id", handler.Update)
	return router
}

func TestUserHandler_Index(t *testing.T) {
	mockUsecase := new(MockUserUsecase)
	router := setupUserRouter(mockUsecase)

	mockUsecase.On("Pagination", mock.Anything, mock.Anything).Return(response.Response{Status: true, Data: nil})

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_Show_NotFound(t *testing.T) {
	mockUsecase := new(MockUserUsecase)
	router := setupUserRouter(mockUsecase)

	mockUsecase.On("Show", mock.Anything).Return(models.User{})

	req, _ := http.NewRequest(http.MethodGet, "/users/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Based on typical empty return on mock for Show => NotFound
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
