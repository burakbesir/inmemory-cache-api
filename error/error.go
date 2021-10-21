package error

import (
	"encoding/json"
)

const (
	ValidationError = "ValidationError"
	NotFoundError   = "NotFoundError"
)

type ErrorResponse struct {
	StatusCode       int    `json:"statusCode"`
	ErrorName        string `json:"errorName"`
	ErrorDescription string `json:"errorDescription"`
}

func (e *ErrorResponse) Error() string {
	errorResponse, _ := json.Marshal(e)
	return string(errorResponse)
}

func CreateNotFoundError(err error) error {
	return &ErrorResponse{
		StatusCode:       404,
		ErrorName:        NotFoundError,
		ErrorDescription: err.Error(),
	}
}

func CreateValidationError(err error) error {
	return &ErrorResponse{
		StatusCode:       400,
		ErrorName:        ValidationError,
		ErrorDescription: err.Error(),
	}
}
