package models

import "time"

type Todo struct {
	ID        string    `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}
