package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `gorm:"unique" json:"username" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required,email"`
	Password string `json:"password"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
	// Role      Role           `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
