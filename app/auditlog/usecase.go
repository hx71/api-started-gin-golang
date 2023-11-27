package auditlog

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/hx71/api-started-gin-golang/response"
)

type Usecase interface {
	Create(req dto.AuditLogCreateValidation) error
	Show(id string) models.AuditLog
	Delete(model models.AuditLog) error
	Pagination(ctx *gin.Context, pagination *response.Pagination) response.Response
}
