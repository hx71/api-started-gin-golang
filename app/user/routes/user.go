package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/user"
	"github.com/hx71/api-started-gin-golang/app/user/handler"
)

func UserHTTPHandler(rg *gin.RouterGroup, usecases user.Usecase) {
	handlers := handler.NewUserHandler(usecases)
	r := rg.Group("/users")
	{
		r.GET("", handlers.Index)
		r.POST("", handlers.Create)
		r.GET("/:id", handlers.Show)
		r.PUT("/:id", handlers.Update)
		r.DELETE("/:id", handlers.Delete)
	}

}
