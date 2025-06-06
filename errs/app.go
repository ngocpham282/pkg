package errs

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	BaseError error  `json:"error,omitempty"`
	Details   any    `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	msg := fmt.Sprintf("%d: %s", e.Code, e.Message)
	if e.BaseError != nil {
		msg += fmt.Sprintf(", baseError: %v", e.BaseError)
	}
	if e.Details != nil {
		msg += fmt.Sprintf(", details: %v", e.Details)
	}
	return msg
}

func (e *AppError) Unwrap() error {
	return e.BaseError
}

func NewAppError(code int) *AppError {
	message, ok := ErrMapping[code]
	if !ok {
		message = ErrMsgInternalError
	}
	return &AppError{
		Code:      code,
		Message:   message,
		BaseError: errors.New(message),
	}
}

func NewAppErrorWithMsg(code int, message string) *AppError {
	return &AppError{
		Code:      code,
		Message:   message,
		BaseError: errors.New(message),
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
