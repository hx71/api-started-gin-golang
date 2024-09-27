package engine

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hx71/api-started-gin-golang/app/auditlog"
	"github.com/hx71/api-started-gin-golang/app/auth"
	"github.com/hx71/api-started-gin-golang/app/jwtauth"
	"github.com/hx71/api-started-gin-golang/app/menu"
	"github.com/hx71/api-started-gin-golang/app/role"
	"github.com/hx71/api-started-gin-golang/app/user"
	"github.com/hx71/api-started-gin-golang/app/usermenu"
	"github.com/hx71/api-started-gin-golang/middleware"

	eAuth "github.com/hx71/api-started-gin-golang/app/auth/routes"
	uAuth "github.com/hx71/api-started-gin-golang/app/auth/usecase"

	rAuditlog "github.com/hx71/api-started-gin-golang/app/auditlog/repository"
	eAuditlog "github.com/hx71/api-started-gin-golang/app/auditlog/routes"
	uAuditlog "github.com/hx71/api-started-gin-golang/app/auditlog/usecase"

	rMenu "github.com/hx71/api-started-gin-golang/app/menu/repository"
	eMenu "github.com/hx71/api-started-gin-golang/app/menu/routes"
	uMenu "github.com/hx71/api-started-gin-golang/app/menu/usecase"

	rRole "github.com/hx71/api-started-gin-golang/app/role/repository"
	eRole "github.com/hx71/api-started-gin-golang/app/role/routes"
	uRole "github.com/hx71/api-started-gin-golang/app/role/usecase"

	rUser "github.com/hx71/api-started-gin-golang/app/user/repository"
	eUser "github.com/hx71/api-started-gin-golang/app/user/routes"
	uUser "github.com/hx71/api-started-gin-golang/app/user/usecase"

	rUserMenu "github.com/hx71/api-started-gin-golang/app/usermenu/repository"
	eUserMenu "github.com/hx71/api-started-gin-golang/app/usermenu/routes"
	uUserMenu "github.com/hx71/api-started-gin-golang/app/usermenu/usecase"

	"github.com/hx71/api-started-gin-golang/config"
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

	// repository
	jwtAuth jwtauth.JWTService = jwtauth.NewJWTService()

	auditlogs auditlog.Repository = rAuditlog.NewAuditLogRepository(db)
	menus     menu.Repository     = rMenu.NewMenuRepository(db)
	roles     role.Repository     = rRole.NewRoleRepository(db)
	users     user.Repository     = rUser.NewUserRepository(db)
	usermenus usermenu.Repository = rUserMenu.NewUserMenuRepository(db)

	// usecase
	authUsecase     auth.Usecase     = uAuth.NewAuthUsecase(users)
	auditlogUsecase auditlog.Usecase = uAuditlog.NewAuditLogUsecase(auditlogs)
	menuUsecase     menu.Usecase     = uMenu.NewMenuUsecase(menus)
	roleUsecase     role.Usecase     = uRole.NewRoleUsecase(roles)
	userUsecase     user.Usecase     = uUser.NewUserUsecase(users)
	userMenuUsecase usermenu.Usecase = uUserMenu.NewUserMenuUsecase(usermenus)
)

func SetupRouter() *gin.Engine {
	// defer config.CloseConnection(dbx)

	// Gin instance
	r := gin.Default()
	// if !envConfig.Debug {
	// 	gin.SetMode(gin.ReleaseMode)
	// }
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	//Logging
	// r.Use(helpers.LoggerToFile())

	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	v1 := r.Group("api/v1")
	{
		v1.GET("/version", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api version 1.0.0",
			})
		})

		// // create log file
		// currentTime := time.Now()
		// crnTime := currentTime.Format("01-02-2006")
		// fileLog := "log-file-" + crnTime + ".log"
		// _, err := os.OpenFile("logging/"+fileLog, os.O_RDONLY, 0644)
		// if err != nil {
		// 	os.OpenFile("logging/"+fileLog, os.O_CREATE, 0644)
		// }

		// audit logs
		eAuth.AuthHTTPHandler(v1, authUsecase, jwtAuth)

		routes := v1.Group("/", middleware.AuthorizeJWT(jwtAuth))
		{
			// audit logs
			eAuditlog.AuditLogHTTPHandler(routes, auditlogUsecase)
			// menus
			eMenu.MenuHTTPHandler(routes, menuUsecase)
			// roles
			eRole.RoleHTTPHandler(routes, roleUsecase)
			// user
			eUser.UserHTTPHandler(routes, userUsecase)
			// user menus
			eUserMenu.UserMenuHTTPHandler(routes, userMenuUsecase)

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
