package errors

import (
	"errors"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

type Error struct {
	Code    int    `json:"code"`
	Error 	string `json:"error"`
	Err     error  `json:"err,omitempty"`
}


// Unwrap returns the inner error.
func (e *Error) Unwrap() error { return e.Err }
func (e *Error) Original() error { return e.Err }

// ---- Core Constructors ----

func New(code int, errMsg string, err ...error) *Error {
	var underlying error
	if len(err) > 0 {
		underlying = err[0]
	}
	return &Error{Code: code, Error: errMsg, Err: underlying}
}


// Internal Server Error
func Wrap(msg string, err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{Code: http.StatusInternalServerError, Error: msg, Err: err}
}



// ---- Common Helpers ----

func BadRequest(msg string) *Error {
	return &Error{Code: http.StatusBadRequest, Error: msg}
}

func Unauthorized(msg string) *Error {
	return &Error{Code: http.StatusUnauthorized, Error: msg}
}

func Forbidden(msg string) *Error {
	return &Error{Code: http.StatusForbidden, Error: msg}
}

func NotFound(msg string) *Error {
	return &Error{Code: http.StatusNotFound, Error: msg}
}

func Conflict(msg string) *Error {
	return &Error{Code: http.StatusConflict, Error: msg}
}

func Internal(err error) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Error: "Internal Server Error",
		Err:     err,
	}
}

func UnprocessableEntity(msg string) *Error {
	return &Error{Code: http.StatusUnprocessableEntity, Error: msg}
}

func ServiceUnavailable(msg string) *Error {
	return &Error{Code: http.StatusServiceUnavailable, Error: msg}
}

// ---- Utilities ----
func IsPgConflict(err error) bool {
    var pgErr *pq.Error
    if errors.As(err, &pgErr) {
        return pgErr.Code == "23505"
    }
    // sometimes GORM wraps the pq.Error in a string, fallback
    return strings.Contains(err.Error(), "duplicate key value")
}