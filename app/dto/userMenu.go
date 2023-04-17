package dto

// Create validation is a model that used by client when POST
type UserMenuCreateValidation struct {
	// ID     string `json:"id" form:"id" binding:"required,omitempty,uuid"`
	RoleID   string `json:"role_id" form:"role_id" binding:"required"`
	MenuID   string `json:"menu_id" form:"menu_id" binding:"required"`
	IsRead   bool   `json:"is_read" form:"is_read"`
	IsCreate bool   `json:"is_create" form:"is_create"`
	IsUpdate bool   `json:"is_update" form:"is_update"`
	IsDelete bool   `json:"is_delete" form:"is_delete"`
	IsReport bool   `json:"is_report" form:"is_report"`
}
