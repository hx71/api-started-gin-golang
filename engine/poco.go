package engine

type ResponseStatus struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type ResponseSuccess struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Role struct {
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

type Menu struct {
	MainMenu  string `json:"main_menu" form:"main_menu" binding:"required"`
	Parent    uint8  `json:"parent" form:"parent"`
	Name      string `json:"name" form:"name" binding:"required"`
	Icon      string `json:"icon" form:"icon" binding:"required"`
	Url       string `json:"url" form:"url" binding:"required"`
	Index     uint16 `json:"index" form:"index" binding:"required"`
	Sort      uint8  `json:"sort" form:"sort" binding:"required"`
	IsActive  bool   `json:"is_active" form:"is_active"`
	SubParent bool   `json:"sub_parent" form:"sub_parent"`
}
