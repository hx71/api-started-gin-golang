package usermenu

import (
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Repository interface {
	Create(model models.UserMenu) error
	Show(id string) models.UserMenus
	Update(model models.UserMenu) error
	Delete(model models.UserMenus) error
	Pagination(*helpers.Pagination) (response.RepositoryResult, int)
}
