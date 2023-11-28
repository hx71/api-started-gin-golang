package role

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Usecase interface {
	Create(req dto.RoleCreateValidation) error
	Show(id string) models.Role
	Update(req dto.RoleCreateValidation) error
	Delete(model models.Role) error
	Pagination(ctx *gin.Context, pagination *response.Pagination) response.Response
}
