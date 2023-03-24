package protocols

import (
	"errors"
	"fmt"
)

type AppError struct {
	StatusCode int
	Err        error
}

func (r *AppError) Error() string {
	return fmt.Sprintf("[STATUS %d]: %v", r.StatusCode, r.Err)
}

func NewAppError(message string, statusCode int) *AppError {
	return &AppError{
		Err:        errors.New(message),
		StatusCode: statusCode,
	}
}

type CryptoCompareError struct {
	Message        string `json:"Message"`
	ParamWithError string `json:"ParamWithError"`
	Response       string `json:"Response"`
	Type           uint32 `json:"Type"`
}

func (e *CryptoCompareError) Error() string {
	message := fmt.Sprintf("[CryptoCompareError]: %s", e.Message)
	return message
}
