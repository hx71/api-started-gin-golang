package dto

// Create validation is a model that used by client when POST
type MenuCreateValidation struct {
	ID        string `json:"id" form:"id" binding:"required,omitempty,uuid"`
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
