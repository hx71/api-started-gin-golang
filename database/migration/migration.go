package migration

import (
	"fmt"
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

	if exist := db_conn.Migrator().HasTable("users"); !exist {
		err := db_conn.Migrator().CreateTable(&models.User{})
		if err == nil {
			fmt.Println("success migrate table users")
		}
	} else {
		fmt.Println("table users already exist")
	}

	if exist := db_conn.Migrator().HasTable("roles"); !exist {
		err := db_conn.Migrator().CreateTable(&models.Roles{})
		if err == nil {
			fmt.Println("success migrate table roles")
		}
	} else {
		fmt.Println("table roles already exist")
	}

	if exist := db_conn.Migrator().HasTable("menus"); !exist {
		err := db_conn.Migrator().CreateTable(&models.Menu{})
		if err == nil {
			fmt.Println("success migrate table menus")
		}
	} else {
		fmt.Println("table menus already exist")
	}

	if exist := db_conn.Migrator().HasTable("user_menus"); !exist {
		err := db_conn.Migrator().CreateTable(&models.UserMenu{})
		if err == nil {
			fmt.Println("success migrate table user_menus")
		}
	} else {
		fmt.Println("table user_menus already exist")
	}

	if exist := db_conn.Migrator().HasTable("audit_logs"); !exist {
		err := db_conn.Migrator().CreateTable(&models.AuditLog{})
		if err == nil {
			fmt.Println("success migrate table audit_logs")
		}
	} else {
		fmt.Println("table audit_logs already exist")
	}

	fmt.Println("success migrate all tables")
	// db_conn.AutoMigrate(&models.UserMenu{})
	// db_conn.AutoMigrate(&models.AuditLog{})
}
