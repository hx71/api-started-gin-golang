package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Usecase interface {
	Create(req dto.MenuCreateValidation) error
	Show(id string) models.Menus
	Update(req dto.MenuCreateValidation) error
	Delete(id string) error
	Pagination(ctx *gin.Context, pagination *helpers.Pagination) response.Response
}
