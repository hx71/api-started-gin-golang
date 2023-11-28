package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/auth"
	"github.com/hx71/api-started-gin-golang/app/auth/handler"
	"github.com/hx71/api-started-gin-golang/app/jwtauth"
	"github.com/hx71/api-started-gin-golang/middleware"
)

func AuthHTTPHandler(rg *gin.RouterGroup, usecase auth.Usecase, jwts jwtauth.JWTService) {
	handlers := handler.NewAuthHandler(usecase)
	r := rg.Group("/auth")
	{
		r.POST("/login", handlers.Login)
		r.POST("/register", handlers.Register)
		r.GET("/logout", middleware.AuthorizeJWT(jwts), handlers.Logout)
	}
}
