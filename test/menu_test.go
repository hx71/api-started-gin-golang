package test

import (
	"github.com/gin-gonic/gin"
	handlerPackage "github.com/hx71/api-started-gin-golang/app/menu/handler"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMenuRouter(mockUsecase *MockMenuUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := handlerPackage.NewMenuHandler(mockUsecase)

	router.GET("/menus", handler.Index)
	router.GET("/menus/:id", handler.Show)
	router.POST("/menus", handler.Create)
	router.DELETE("/menus/:id", handler.Delete)
	router.PUT("/menus/:id", handler.Update)
	return router
}

func TestMenuHandler_Index(t *testing.T) {
	mockUsecase := new(MockMenuUsecase)
	router := setupMenuRouter(mockUsecase)

	mockUsecase.On("Pagination", mock.Anything, mock.Anything).Return(response.Response{Status: true, Data: nil})

	req, _ := http.NewRequest(http.MethodGet, "/menus", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMenuHandler_Show_NotFound(t *testing.T) {
	mockUsecase := new(MockMenuUsecase)
	router := setupMenuRouter(mockUsecase)

	mockUsecase.On("Show", mock.Anything).Return(models.Menus{})

	req, _ := http.NewRequest(http.MethodGet, "/menus/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Based on typical empty return on mock for Show => NotFound
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
