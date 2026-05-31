package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/hx71/api-started-gin-golang/app/dto"
	roleRepo "github.com/hx71/api-started-gin-golang/app/role/repository"
	roleUsecase "github.com/hx71/api-started-gin-golang/app/role/usecase"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/models"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func setupHandler() *gin.Engine {
	// db := setupTestDB()
	db := config.SetupConnection()

	repo := roleRepo.NewRoleRepository(db)
	usecase := roleUsecase.NewRoleUsecase(repo)
	handler := NewRoleHandler(usecase)

	router := setupRouter()
	router.POST("/roles", handler.Create)
	router.GET("/roles/:id", handler.Show)
	router.PUT("/roles/:id", handler.Update)
	router.DELETE("/roles/:id", handler.Delete)

	return router
}

func TestRoleHandler_Create_Postgres(t *testing.T) {
	router := setupHandler()

	payload := dto.RoleCreateValidation{
		Name: "Admin",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/roles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestRoleHandler_Show_Postgres(t *testing.T) {
	db := setupTestDB()

	roleID := uuid.NewString()
	db.Create(&models.Role{
		ID:   roleID,
		Name: "Admin",
	})

	repo := roleRepo.NewRoleRepository(db)
	usecase := roleUsecase.NewRoleUsecase(repo)
	handler := NewRoleHandler(usecase)

	router := setupRouter()
	router.GET("/roles/:id", handler.Show)

	req, _ := http.NewRequest(http.MethodGet, "/roles/"+roleID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRoleHandler_Update_Postgres(t *testing.T) {
	db := setupTestDB()

	roleID := uuid.NewString()
	db.Create(&models.Role{
		ID:   roleID,
		Name: "Admin",
	})

	repo := roleRepo.NewRoleRepository(db)
	usecase := roleUsecase.NewRoleUsecase(repo)
	handler := NewRoleHandler(usecase)

	router := setupRouter()
	router.PUT("/roles/:id", handler.Update)

	payload := dto.RoleCreateValidation{
		Name: "Super Admin",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPut, "/roles/"+roleID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestRoleHandler_Delete_Postgres(t *testing.T) {
	db := setupTestDB()

	roleID := uuid.NewString()
	db.Create(&models.Role{
		ID:   roleID,
		Name: "Admin",
	})

	repo := roleRepo.NewRoleRepository(db)
	usecase := roleUsecase.NewRoleUsecase(repo)
	handler := NewRoleHandler(usecase)

	router := setupRouter()
	router.DELETE("/roles/:id", handler.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/roles/"+roleID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRoleHandler_Show_NotFound_Postgres(t *testing.T) {
	router := setupHandler()

	req, _ := http.NewRequest(http.MethodGet, "/roles/invalid-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
