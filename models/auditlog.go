package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID          string `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	UserID      string `gorm:"not null;type:uuid;index" json:"user_id,omitempty"`
	IPAddress   string `gorm:"not null" json:"ip_address"`
	ServiceName string `gorm:"not null" json:"service_name"` // user
	MethodName  string `gorm:"not null" json:"method_name"`  // create user
	// ServiceType string         `gorm:"not null" json:"service_type"` // user
	// Action      string         `gorm:"not null" json:"action"`       // create user
	Level     string         `gorm:"not null" json:"level"`    // info or error
	Metadata  string         `gorm:"not null" json:"metadata"` // response
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
