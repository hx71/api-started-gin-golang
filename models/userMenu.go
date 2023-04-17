package models

import (
	"time"

	"gorm.io/gorm"
)

type UserMenu struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	RoleID    string         `json:"role_id"`
	MenuID    string         `json:"menu_id"`
	IsRead    bool           `json:"is_read"`
	IsCreate  bool           `json:"is_create"`
	IsUpdate  bool           `json:"is_update"`
	IsDelete  bool           `json:"is_delete"`
	IsReport  bool           `json:"is_report"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (UserMenus) TableName() string { return "user_menus" }

type UserMenus struct {
	ID        string         `gorm:"primary_key" json:"id"`
	RoleID    string         `json:"role_id"`
	MenuID    string         `json:"menu_id"`
	IsRead    bool           `json:"is_read"`
	IsCreate  bool           `json:"is_create"`
	IsUpdate  bool           `json:"is_update"`
	IsDelete  bool           `json:"is_delete"`
	IsReport  bool           `json:"is_report"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
