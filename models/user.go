package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	Name      string         `json:"name" binding:"required"`
	Username  string         `gorm:"unique" json:"username" binding:"required"`
	Email     string         `gorm:"unique" json:"email" binding:"required,email"`
	Password  string         `json:"password"`
	Token     string         `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
