package config

type ResError struct {
	FailedProcess string
}

func NewResError() *ResError {
	return &ResError{
		FailedProcess: "Failed to process request",
	}
}
