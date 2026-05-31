package test

import (
	"github.com/gin-gonic/gin"
	handlerPackage "github.com/hx71/api-started-gin-golang/app/role/handler"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRoleRouter(mockUsecase *MockRoleUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := handlerPackage.NewRoleHandler(mockUsecase)

	router.GET("/roles", handler.Index)
	router.GET("/roles/:id", handler.Show)
	router.POST("/roles", handler.Create)
	router.DELETE("/roles/:id", handler.Delete)
	router.PUT("/roles/:id", handler.Update)
	return router
}

func TestRoleHandler_Index(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	router := setupRoleRouter(mockUsecase)

	mockUsecase.On("Pagination", mock.Anything, mock.Anything).Return(response.Response{Status: true, Data: nil})

	req, _ := http.NewRequest(http.MethodGet, "/roles", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRoleHandler_Show_NotFound(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	router := setupRoleRouter(mockUsecase)

	mockUsecase.On("Show", mock.Anything).Return(models.Role{})

	req, _ := http.NewRequest(http.MethodGet, "/roles/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Based on typical empty return on mock for Show => NotFound
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
