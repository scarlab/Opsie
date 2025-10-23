package types

import (
	"net/http"
	"opsie/pkg/errors"
)

// HandlerFunc is a function that handles HTTP requests.
// This is a simple shorthand to define easier to read functions.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) *errors.Error

// Middleware is a special type that handles HandleFuncs.
type Middleware func(HandlerFunc) HandlerFunc
