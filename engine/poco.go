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

type Role struct {
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}
