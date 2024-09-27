package response

import "strings"

type Result struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ResultSuccess returns a successful Result with the given message.
func ResultSuccess(message string) Result {
	return Result{
		Status:  true,
		Message: message,
	}
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseSuccess creates a success response with a message and data.
func ResponseSuccess(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

type ResError struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// ResponseError returns a strongly typed ResError struct with the given message and error.
func ResponseError(message string, err string) ResError {
	return ResError{
		Status:  false,
		Message: message,
		Errors:  strings.Split(err, "\n"),
	}
}
