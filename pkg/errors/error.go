package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}
// Unwrap returns the inner error.
func (e *Error) Unwrap() error {
	return e.Err
}

// ---- Core Constructors ----

func New(code int, msg string, err ...error) *Error {
	var underlying error
	if len(err) > 0 {
		underlying = err[0]
	}
	return &Error{Code: code, Message: msg, Err: underlying}
}


// Internal Server Error
func Wrap(msg string, err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{Code: http.StatusInternalServerError, Message: msg, Err: err}
}



// ---- Common Helpers ----

func BadRequest(msg string) *Error {
	return &Error{Code: http.StatusBadRequest, Message: msg}
}

func Unauthorized(msg string) *Error {
	return &Error{Code: http.StatusUnauthorized, Message: msg}
}

func Forbidden(msg string) *Error {
	return &Error{Code: http.StatusForbidden, Message: msg}
}

func NotFound(msg string) *Error {
	return &Error{Code: http.StatusNotFound, Message: msg}
}

func Conflict(msg string) *Error {
	return &Error{Code: http.StatusConflict, Message: msg}
}

func Internal(err error) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Err:     err,
	}
}

func UnprocessableEntity(msg string) *Error {
	return &Error{Code: http.StatusUnprocessableEntity, Message: msg}
}

func ServiceUnavailable(msg string) *Error {
	return &Error{Code: http.StatusServiceUnavailable, Message: msg}
}

// ---- Utilities ----

// Is checks if error matches a specific code.
func Is(err error, code int) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Code == code
	}
	return false
}


