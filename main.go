package main

import (
	"log"
	"os"
	"srp-golang/app/controllers"
	"srp-golang/config"
	"srp-golang/middleware"
	"srp-golang/repository"
	"srp-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	db             *gorm.DB                   = config.SetupConnection()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	authService    service.AuthService        = service.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	jwtService     service.JWTService         = service.NewJWTService(userRepository)
)

func main() {
	defer config.CloseConnection(db)

	r := gin.Default()

	auth := r.Group("api/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	rt := r.Group("api", middleware.AuthorizeJWT(jwtService))
	// rt := r.Group("api")
	{
		rt.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run(":" + os.Getenv("APP_PORT"))
}
