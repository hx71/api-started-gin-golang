package dto

// Create validation is a model that used by client when POST
type RoleCreateValidation struct {
	ID   string `json:"id" form:"id" binding:"required,omitempty,uuid"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}
