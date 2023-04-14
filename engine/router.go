package engine

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/controllers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/repository"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/app/service"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/config"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/helpers"
	"github.com/hasrulrhul/service-repository-pattern-gin-golang/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	_ "github.com/hasrulrhul/service-repository-pattern-gin-golang/docs/swagger"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var (
	db *gorm.DB = config.SetupConnection()

	jwtService service.JWTService = service.NewJWTService()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	roleRepository repository.RoleRepository = repository.NewRoleRepository(db)
	menuRepository repository.MenuRepository = repository.NewMenuRepository(db)

	authService service.AuthService = service.NewAuthService(userRepository)
	userService service.UserService = service.NewUserService(userRepository)
	roleService service.RoleService = service.NewRoleService(roleRepository)
	menuService service.MenuService = service.NewMenuService(menuRepository)

	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
	roleController controllers.RoleController = controllers.NewRoleController(roleService, jwtService)
	menuController controllers.MenuController = controllers.NewMenuController(menuService, jwtService)
)

func SetupRouter() *gin.Engine {
	// defer config.CloseConnection(db)

	// Gin instance
	r := gin.Default()
	// if !envConfig.Debug {
	// 	gin.SetMode(gin.ReleaseMode)
	// }
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	//Logging
	r.Use(helpers.LoggerToFile())

	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	v1 := r.Group("api/v1")
	{
		currentTime := time.Now()
		crnTime := currentTime.Format("01-02-2006")
		// log file
		fileLog := "log-file-" + crnTime + ".log"

		_, err := os.OpenFile("logging/"+fileLog, os.O_RDONLY, 0644)
		if err != nil {
			os.OpenFile("logging/"+fileLog, os.O_CREATE, 0644)
		}

		v1.GET("/version", authController.Version)

		auth := v1.Group("auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
			auth.GET("/logout", middleware.AuthorizeJWT(jwtService), authController.Logout)
		}

		routes := v1.Group("/")
		// routes := v1.Group("/", middleware.AuthorizeJWT(jwtService))
		{
			users := routes.Group("/users")
			{
				users.GET("", userController.Index)
				users.POST("", userController.Create)
				users.GET("/:id", userController.Show)
				users.PUT("/:id", userController.Update)
				users.DELETE("/:id", userController.Delete)
			}

			role := routes.Group("/roles")
			{
				role.GET("", roleController.Index)
				role.POST("", roleController.Create)
				role.GET("/:id", roleController.Show)
				role.PUT("/:id", roleController.Update)
				role.DELETE("/:id", roleController.Delete)
			}

			menu := routes.Group("/menus")
			{
				menu.GET("", menuController.Index)
				menu.POST("", menuController.Create)
				menu.GET("/:id", menuController.Show)
				menu.PUT("/:id", menuController.Update)
				menu.DELETE("/:id", menuController.Delete)
			}
		}
	}
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			//c.Next()
			return
		}
		c.Next()
	}
}
