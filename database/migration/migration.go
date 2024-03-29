package migration

import (
	"log"

	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/models"
	"gorm.io/gorm"
)

var (
	db_conn *gorm.DB = config.SetupConnection()
)

func RunMigrations() {
	if db_conn.Error != nil {
		log.Fatalln(db_conn.Error.Error())
	}

	db_conn.AutoMigrate(&models.User{})
	db_conn.AutoMigrate(&models.Role{})
	db_conn.AutoMigrate(&models.Menu{})
	db_conn.AutoMigrate(&models.UserMenu{})
	db_conn.AutoMigrate(&models.AuditLog{})
}
