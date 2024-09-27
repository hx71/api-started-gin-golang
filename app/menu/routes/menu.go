package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/menu"
	"github.com/hx71/api-started-gin-golang/app/menu/handler"
)

func MenuHTTPHandler(rg *gin.RouterGroup, usecases menu.Usecase) {
	handlers := handler.NewMenuHandler(usecases)
	r := rg.Group("/menus")
	{
		r.GET("", handlers.Index)
		r.POST("", handlers.Create)
		r.GET("/:id", handlers.Show)
		r.PUT("/:id", handlers.Update)
		r.DELETE("/:id", handlers.Delete)
	}
}
