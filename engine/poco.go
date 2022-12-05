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
