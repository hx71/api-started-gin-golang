package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/auditlog"
	"github.com/hx71/api-started-gin-golang/app/auditlog/handler"
)

func AuditLogHTTPHandler(rg *gin.RouterGroup, usecases auditlog.Usecase) {
	handlers := handler.NewAuditLogHandler(usecases)
	r := rg.Group("/audit-log")
	{
		r.GET("", handlers.Index)
		r.POST("", handlers.Create)
		r.GET("/:id", handlers.Show)
		r.DELETE("/:id", handlers.Delete)
	}
}
