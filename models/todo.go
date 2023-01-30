package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	Name      string         `json:"name"`
	UserID    string         `json:"user_id"  foreigen_key:"true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Todos) TableName() string { return "todos" }

type Todos struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	Name      string         `json:"name"`
	NameUser  string         `json:"name_user"`
	UserID    string         `json:"user_id"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
