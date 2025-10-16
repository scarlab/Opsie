package bolt

import "fmt"



type Error struct {
	Code    int 
	Message string
}

// Implement error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Helper constructors
func NewError(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}


