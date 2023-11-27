package engine

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/auditlog"
	auditlogRepo "github.com/hx71/api-started-gin-golang/app/auditlog/repository"
	rAuditlog "github.com/hx71/api-started-gin-golang/app/auditlog/routes"
	"github.com/hx71/api-started-gin-golang/app/auditlog/usecase"
	"github.com/hx71/api-started-gin-golang/app/controllers"
	"github.com/hx71/api-started-gin-golang/app/repository"
	"github.com/hx71/api-started-gin-golang/app/service"
	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/helpers"
	"github.com/hx71/api-started-gin-golang/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	_ "github.com/hx71/api-started-gin-golang/docs/swagger"
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

	users     repository.UserRepository     = repository.NewUserRepository(db)
	roles     repository.RoleRepository     = repository.NewRoleRepository(db)
	menus     repository.MenuRepository     = repository.NewMenuRepository(db)
	userMenus repository.UserMenuRepository = repository.NewUserMenuRepository(db)

	authService     service.AuthService     = service.NewAuthService(users)
	userService     service.UserService     = service.NewUserService(users)
	roleService     service.RoleService     = service.NewRoleService(roles)
	menuService     service.MenuService     = service.NewMenuService(menus)
	userMenuService service.UserMenuService = service.NewUserMenuService(userMenus)

	authController     controllers.AuthController     = controllers.NewAuthController(authService, jwtService)
	userController     controllers.UserController     = controllers.NewUserController(userService, jwtService)
	roleController     controllers.RoleController     = controllers.NewRoleController(roleService, jwtService)
	menuController     controllers.MenuController     = controllers.NewMenuController(menuService, jwtService)
	userMenuController controllers.UserMenuController = controllers.NewUserMenuController(userMenuService, jwtService)

	auditlogs       auditlog.Repository = auditlogRepo.NewAuditLogRepository(db)
	auditlogUsecase auditlog.Usecase    = usecase.NewAuditLogUsecase(auditlogs)
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

			// audit logs
			rAuditlog.AuditLogHTTPHandler(routes, auditlogUsecase)

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

			userMenu := routes.Group("/user-menus")
			{
				userMenu.GET("", userMenuController.Index)
				userMenu.POST("", userMenuController.Create)
				userMenu.GET("/:id", userMenuController.Show)
				userMenu.PUT("/:id", userMenuController.Update)
				userMenu.DELETE("/:id", userMenuController.Delete)
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
