package bolt

import (
	"log"
	"net/http"
	"opsie/config"
	"time"
)

// HandlerFunc is a function that handles HTTP requests.
// This is a simple shorthand to define easier to read functions.
type THandlerFunc func(w http.ResponseWriter, r *http.Request)

// Middleware is a special type that handles HandleFuncs.
type TMiddleware func(THandlerFunc) THandlerFunc




// Handle handles the middlewares.
// It executes the middlewares in the order presented and finishes by calling the final handler.
func Middleware(final THandlerFunc, middlewares ...TMiddleware) THandlerFunc {
	if final == nil {
		panic("no final handler")
		// Or return a default handler.
	}
	
	if config.IsDev {
		middlewares = append(middlewares, logger)
	}

	// Execute the middleware in the same order and return the final func.
	// This is a confusing and tricky construct :)
	// We need to use the reverse order since we are chaining inwards.
	for i := len(middlewares) - 1; i >= 0; i-- {
		final = middlewares[i](final) // mw1(mw2(mw3(final)))
	}
	return final
}




// logger measures request time, status, and response size.
func logger(next THandlerFunc) THandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the ResponseWriter
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

		next(lrw, r) // call the next handler

		duration := time.Since(start)
		log.Printf("%s %s %d %v %d", r.Method, r.RequestURI, lrw.status, duration, lrw.size)
	}
}

// loggingResponseWriter captures status code and response size
type loggingResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.size += size
	return size, err
}
