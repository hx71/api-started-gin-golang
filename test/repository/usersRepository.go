package repository

import (
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/models"
)

type UsersRepository interface {
	FindByID(id string) *models.Users
	// Show(id string) *models.Users
}
