package usermenu

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Usecase interface {
	Create(req []dto.UserMenuCreateValidation) error
	Show(id string) models.UserMenus
	Update(id string, req dto.UserMenuCreateValidation) error
	Delete(model models.UserMenus) error
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}
