package role

import (
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Repository interface {
	Create(model models.Role) error
	Show(id string) models.Role
	Update(req models.Role) error
	Delete(model models.Role) error
	Pagination(*response.Pagination) (response.RepositoryResult, int)
}
