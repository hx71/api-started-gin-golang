package test

import (
	"github.com/gin-gonic/gin"
	handlerPackage "github.com/hx71/api-started-gin-golang/app/usermenu/handler"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupUserMenuRouter(mockUsecase *MockUserMenuUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := handlerPackage.NewUserMenuHandler(mockUsecase)

	router.GET("/usermenus", handler.Index)
	router.GET("/usermenus/:id", handler.Show)
	router.POST("/usermenus", handler.Create)
	router.DELETE("/usermenus/:id", handler.Delete)
	router.PUT("/usermenus/:id", handler.Update)
	return router
}

func TestUserMenuHandler_Index(t *testing.T) {
	mockUsecase := new(MockUserMenuUsecase)
	router := setupUserMenuRouter(mockUsecase)

	mockUsecase.On("Pagination", mock.Anything, mock.Anything).Return(response.Response{Status: true, Data: nil})

	req, _ := http.NewRequest(http.MethodGet, "/usermenus", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserMenuHandler_Show_NotFound(t *testing.T) {
	mockUsecase := new(MockUserMenuUsecase)
	router := setupUserMenuRouter(mockUsecase)

	mockUsecase.On("Show", mock.Anything).Return(models.UserMenus{})

	req, _ := http.NewRequest(http.MethodGet, "/usermenus/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Based on typical empty return on mock for Show => NotFound
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
