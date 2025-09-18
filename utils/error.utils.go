package utils

import (
	"fmt"
)

type CsError struct {
	Code    int 
	Message string
}

// Implement error interface
func (e *CsError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Helper constructors
func NewCsError(code int, msg string) *CsError {
	return &CsError{Code: code, Message: msg}
}


