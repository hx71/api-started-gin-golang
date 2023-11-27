package auditlog

import (
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Repository interface {
	Create(model models.AuditLog) error
	Show(id string) models.AuditLog
	Delete(model models.AuditLog) error
	Pagination(*response.Pagination) (response.RepositoryResult, int)
}
