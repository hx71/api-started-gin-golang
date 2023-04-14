package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID        string         `gorm:"primary_key, not null" json:"id"  binding:"omitempty,uuid"`
	MainMenu  string         `json:"main_menu" gorm:"type:varchar(50)"`
	Parent    uint8          `json:"parent"`
	Name      string         `json:"name" gorm:"type:varchar(100)"`
	Icon      string         `json:"icon" gorm:"type:varchar(100)"`
	Url       string         `json:"url" gorm:"type:varchar(255)"`
	Index     uint16         `json:"index"`
	Sort      uint8          `json:"sort"`
	IsActive  bool           `json:"is_active"`
	SubParent bool           `json:"sub_parent"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Menus) TableName() string { return "menus" }

type Menus struct {
	ID        string         `gorm:"primary_key" json:"id"`
	MainMenu  string         `json:"main_menu"`
	Parent    uint8          `json:"parent"`
	Name      string         `json:"name"`
	Icon      string         `json:"icon"`
	Url       string         `json:"url"`
	Index     uint16         `json:"index"`
	Sort      uint8          `json:"sort"`
	IsActive  bool           `json:"is_active"`
	SubParent bool           `json:"sub_parent"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
