package apif

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	BaseError error  `json:"error,omitempty"`
	Details   any    `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Error: %v, Details: %v", e.Code, e.Message, e.Error, e.Details)
}

func (e *AppError) Unwrap() error {
	return e.BaseError
}

func NewAppError(code int) *AppError {
	return &AppError{
		Code:    code,
		Message: http.StatusText(code),
	}
}
func NewAppErrorWithMsg(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
func (e *AppError) WithError(err error) *AppError {
	e.BaseError = err
	return e
}
func (e *AppError) WithDetails(details any) *AppError {
	e.Details = details
	return e
}
