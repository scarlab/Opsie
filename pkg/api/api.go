package api

import (
	"net/http"
	"opsie/core/mw"
	"opsie/types"
	"strings"

	"github.com/go-chi/chi/v5"
)

// HTTP Method for Api
type HTTPMethod string

const (
	M_GET     	HTTPMethod 		= 		"GET"
	M_POST    	HTTPMethod 		= 		"POST"
	M_PUT     	HTTPMethod 		= 		"PUT"
	M_PATCH   	HTTPMethod 		= 		"PATCH"
	M_DELETE  	HTTPMethod 		= 		"DELETE"
	M_OPTIONS 	HTTPMethod 		= 		"OPTIONS"
	M_HEAD    	HTTPMethod 		= 		"HEAD"
)

func (k HTTPMethod) ToString() string  {
	return string(k)
}

// Api helper
func Api(r chi.Router, method HTTPMethod, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)

	switch method {
	case M_GET:
		r.Get(checkPattern(pattern), final)
	case M_POST:
		r.Post(checkPattern(pattern), final)
	case M_PUT:
		r.Put(checkPattern(pattern), final)
	case M_PATCH:
		r.Put(checkPattern(pattern), final)
	case M_DELETE:
		r.Delete(checkPattern(pattern), final)
	case M_OPTIONS:
		r.Delete(checkPattern(pattern), final)
	case M_HEAD:
		r.Delete(checkPattern(pattern), final)
	default:
		// fallback to all methods
		r.HandleFunc(checkPattern(pattern), final)
	}
}




// Handle handles the middlewares.
// It executes the middlewares in the order presented and finishes by calling the final handler.
func HandleMiddleware(final types.HandlerFunc, middlewares ...types.Middleware) http.HandlerFunc {
	if final == nil {
		panic("no final handler")
		// Or return a default handler.
	}
	
	// --- Default Middlewares
	// --- 0. HttpLogger
	// --- 1. Recover
	middlewares = append([]types.Middleware{mw.HTTPLogger, mw.Recoverer}, middlewares...) // 1th
		

	// --- Execute the middleware in the same order and return the final func.
	// --- This is a confusing and tricky construct :)
	// --- We need to use the reverse order since we are chaining inwards.
	for i := len(middlewares) - 1; i >= 0; i-- {
		final = middlewares[i](final) // mw1(mw2(mw3(final)))
	}

	// Return http handler func
	return func(w http.ResponseWriter, r *http.Request) {
		_ = final(w, r)
	}
}

// checkPattern sanitizes a route path
func checkPattern(pattern string) string {
	if pattern == "" {
		return "/"
	}

	// Ensure single leading slash
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	// Replace multiple slashes with a single slash
	pattern = strings.ReplaceAll(pattern, "//", "/")

	// Remove trailing slash if not root
	if len(pattern) > 1 && strings.HasSuffix(pattern, "/") {
		pattern = strings.TrimSuffix(pattern, "/")
	}

	return pattern
}

/// ______________________________________________________________________
/// Individual helpers ---------------------------------------------------
/// --- 

func Get(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Get(checkPattern(pattern), final)
}

func Post(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Post(checkPattern(pattern), final)
}

func Put(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Put(checkPattern(pattern), final)
}

func Patch(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Patch(checkPattern(pattern), final)
}

func Delete(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Delete(checkPattern(pattern), final)
}

func Options(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Options(checkPattern(pattern), final)
}

func Head(r chi.Router, pattern string, handler types.HandlerFunc, middlewares ...types.Middleware) {
	final := HandleMiddleware(handler, middlewares...)
	r.Head(checkPattern(pattern), final)
}


/// __________________________________
/// --- Here it is [24.10.25] --------
/// --- Find it
/// --- 