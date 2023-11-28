package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/role"
	"github.com/hx71/api-started-gin-golang/app/role/handler"
)

func RoleHTTPHandler(rg *gin.RouterGroup, usecases role.Usecase) {
	handlers := &handler.RoleHandler{Usecase: usecases}
	r := rg.Group("/roles")
	{
		r.GET("", handlers.Index)
		r.POST("", handlers.Create)
		r.GET("/:id", handlers.Show)
		r.PUT("/:id", handlers.Update)
		r.DELETE("/:id", handlers.Delete)
	}
}
