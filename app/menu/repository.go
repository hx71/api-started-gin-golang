package menu

import (
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Repository interface {
	Create(model models.Menu) error
	Show(id string) models.Menus
	Update(model models.Menu) error
	Delete(id string) error
	Pagination(*helpers.Pagination) (response.RepositoryResult, int)
}
