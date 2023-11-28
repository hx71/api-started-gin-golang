package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Usecase interface {
	Create(model dto.UserCreateValidation) error
	Show(id string) models.User
	Update(model dto.UserUpdateValidation) error
	Delete(model models.User) error
	FindByEmail(email string) bool
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}
