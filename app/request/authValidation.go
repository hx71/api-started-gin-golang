package request

//LoginDTO is a model that used by client when POST from /login url
type LoginValidation struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterValidation struct {
	Username string `json:"username" form:"username" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}
