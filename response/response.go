package response

import "strings"

type Result struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// BuildResponse method is to inject data value to dynamic success response
func ResultSuccess(message string) Result {
	res := Result{
		Status:  true,
		Message: message,
	}
	return res
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// BuildResponse method is to inject data value to dynamic success response
func ResponseSuccess(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}

type ResError struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// BuildErrorResponse method is to inject data value to dynamic failed response
func ResponseError(message string, err string) ResError {
	splittedError := strings.Split(err, "\n")
	res := ResError{
		Status:  false,
		Message: message,
		Errors:  splittedError,
	}
	return res
}
