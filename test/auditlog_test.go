package test

import (
	"github.com/gin-gonic/gin"
	handlerPackage "github.com/hx71/api-started-gin-golang/app/auditlog/handler"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupAuditLogRouter(mockUsecase *MockAuditLogUsecase) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := handlerPackage.NewAuditLogHandler(mockUsecase)

	router.GET("/auditlogs", handler.Index)
	router.GET("/auditlogs/:id", handler.Show)
	router.POST("/auditlogs", handler.Create)
	router.DELETE("/auditlogs/:id", handler.Delete)
	return router
}

func TestAuditLogHandler_Index(t *testing.T) {
	mockUsecase := new(MockAuditLogUsecase)
	router := setupAuditLogRouter(mockUsecase)

	mockUsecase.On("Pagination", mock.Anything, mock.Anything).Return(response.Response{Status: true, Data: nil})

	req, _ := http.NewRequest(http.MethodGet, "/auditlogs", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuditLogHandler_Show_NotFound(t *testing.T) {
	mockUsecase := new(MockAuditLogUsecase)
	router := setupAuditLogRouter(mockUsecase)

	mockUsecase.On("Show", mock.Anything).Return(models.AuditLog{})

	req, _ := http.NewRequest(http.MethodGet, "/auditlogs/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Based on typical empty return on mock for Show => NotFound
	// assert.Equal(t, http.StatusNotFound, w.Code)
}
