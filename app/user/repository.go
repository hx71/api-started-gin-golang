package user

import (
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Repository interface {
	Create(model models.User) error
	Show(id string) models.User
	Update(model models.User) error
	Delete(model models.User) error
	Pagination(*helpers.Pagination) (response.RepositoryResult, int)
	VerifyCredential(email string, password string) interface{}
	FindByEmail(email string) bool
}
