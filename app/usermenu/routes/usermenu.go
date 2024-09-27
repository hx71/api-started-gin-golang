package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/usermenu"
	"github.com/hx71/api-started-gin-golang/app/usermenu/handler"
)

func UserMenuHTTPHandler(rg *gin.RouterGroup, usecases usermenu.Usecase) {
	handlers := handler.NewUserMenuHandler(usecases)
	r := rg.Group("/user-menus")
	{
		r.GET("", handlers.Index)
		r.POST("", handlers.Create)
		r.GET("/:id", handlers.Show)
		r.PUT("/:id", handlers.Update)
		r.DELETE("/:id", handlers.Delete)
	}
}
