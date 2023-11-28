package seeder

import (
	"fmt"
	"time"

	"github.com/hx71/api-started-gin-golang/config"
	"github.com/hx71/api-started-gin-golang/models"
	"gorm.io/gorm"
)

var (
	db_conn *gorm.DB = config.SetupConnection()
)

// Roles ..
func Roles() {

	now := time.Now()

	var role []models.Role

	var roles = models.Role{
		ID:        "00000000-0000-0000-0000-000000000001",
		Code:      "A",
		Name:      "Admin",
		CreatedAt: now,
		UpdatedAt: now,
	}
	role = append(role, roles)

	var roles1 = models.Role{
		ID:        "00000000-0000-0000-0000-000000000002",
		Code:      "U",
		Name:      "User",
		CreatedAt: now,
		UpdatedAt: now,
	}
	role = append(role, roles1)

	for _, v := range role {
		if err := db_conn.Where("code = ?", v.Code).First(&role).Error; err != nil {
			db_conn.Create(&v)
		}
		fmt.Printf("roles %s has been created\n", v.Code)
	}
}

// Menus ..
func Menus() {

	now := time.Now()

	var menu []models.Menu

	var menus = models.Menu{
		ID:        "00000000-0000-0000-0000-000000000001",
		MainMenu:  "Apps",
		Parent:    0,
		Name:      "Roles",
		Icon:      "fa-gear",
		Url:       "roles",
		Index:     100,
		Sort:      1,
		IsActive:  true,
		SubParent: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	menu = append(menu, menus)

	var menus1 = models.Menu{
		ID:        "00000000-0000-0000-0000-000000000002",
		MainMenu:  "Apps",
		Parent:    0,
		Name:      "Menus",
		Icon:      "fa-gear",
		Url:       "menus",
		Index:     100,
		Sort:      2,
		IsActive:  true,
		SubParent: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	menu = append(menu, menus1)

	var menus2 = models.Menu{
		ID:        "00000000-0000-0000-0000-000000000003",
		MainMenu:  "Apps",
		Parent:    0,
		Name:      "Users",
		Icon:      "fa-user",
		Url:       "users",
		Index:     100,
		Sort:      2,
		IsActive:  true,
		SubParent: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	menu = append(menu, menus2)

	for _, v := range menu {
		if err := db_conn.Where("id = ?", v.ID).First(&menu).Error; err != nil {
			db_conn.Create(&v)
		}
		fmt.Printf("menu %s has been created\n", v.Name)
	}
}

// User-Menus ..
func UserMenus() {

	now := time.Now()

	var user_menu []models.UserMenu

	var user_menus = models.UserMenu{
		ID:        "00000000-0000-0000-0000-000000000001",
		RoleID:    "00000000-0000-0000-0000-000000000001",
		MenuID:    "00000000-0000-0000-0000-000000000001",
		IsRead:    true,
		IsCreate:  true,
		IsUpdate:  true,
		IsDelete:  true,
		IsReport:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	user_menu = append(user_menu, user_menus)

	var user_menus1 = models.UserMenu{
		ID:        "00000000-0000-0000-0000-000000000002",
		RoleID:    "00000000-0000-0000-0000-000000000002",
		MenuID:    "00000000-0000-0000-0000-000000000002",
		IsRead:    true,
		IsCreate:  true,
		IsUpdate:  true,
		IsDelete:  true,
		IsReport:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	user_menu = append(user_menu, user_menus1)

	for _, v := range user_menu {
		if err := db_conn.Where("id = ?", v.ID).First(&user_menu).Error; err != nil {
			db_conn.Create(&v)
		}
		fmt.Printf("user-menu %s has been created\n", v.ID)
	}
}

// Users ..
func Users() {

	now := time.Now()

	var user []models.User

	var users = models.User{
		ID:        "00000000-0000-0000-0000-000000000001",
		Name:      "Administrator",
		Username:  "admin",
		Email:     "admin@example.com",
		Password:  "$2a$04$W/5vD3bgWtcmMz7yWXmV2OJUUNz/hMPPnsHUCww8z.gZXlWwTRV7C", // 123
		CreatedAt: now,
		UpdatedAt: now,
	}
	user = append(user, users)

	for _, v := range user {
		if err := db_conn.Where("username = ?", v.Username).First(&user).Error; err != nil {
			db_conn.Create(&v)
		}
		fmt.Printf("users %s has been created\n", v.Name)
	}
}

func RunSeeder() {
	Roles()
	Menus()
	UserMenus()
	Users()
}
