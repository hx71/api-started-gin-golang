package dto

// Create validation is a model that used by client when POST
type TodoCreateValidation struct {
	ID     string `json:"id" form:"id" binding:"required,omitempty,uuid"`
	Name   string `json:"name" form:"name" binding:"required"`
	UserID string `json:"user_id" form:"user_id" binding:"required,omitempty,uuid"`
}
