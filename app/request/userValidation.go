package request

// Create validation is a model that used by client when POST
type UserCreateValidation struct {
	Username string `json:"username" form:"username" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}

//  Update validation is a model that used by client when POST
type UserUpdateValidation struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}
