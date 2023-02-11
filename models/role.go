package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Roles) TableName() string { return "roles" }

type Roles struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
