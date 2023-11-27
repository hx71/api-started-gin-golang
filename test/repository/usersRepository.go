package repository

import (
	"github.com/hx71/api-started-gin-golang/models"
)

type UsersRepository interface {
	Index() *[]models.User
	Create(req models.User) (*models.User, error)
	Show(id string) *models.User
	Update(req models.User) (*models.User, error)
	Delete(req models.User) (*models.User, error)
}
