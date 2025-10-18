package bolt

import (
	"log"
	"net/http"
	"opsie/config"
	"opsie/pkg/errors"
	"reflect"
	"time"
)

// HandlerFunc is a function that handles HTTP requests.
// This is a simple shorthand to define easier to read functions.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) *errors.Error

// Middleware is a special type that handles HandleFuncs.
type Middleware func(HandlerFunc) HandlerFunc




// Handle handles the middlewares.
// It executes the middlewares in the order presented and finishes by calling the final handler.
func HandleMiddleware(final HandlerFunc, middlewares ...Middleware) HandlerFunc {
	if final == nil {
		panic("no final handler")
		// Or return a default handler.
	}
	
	// Add Error Handler to the chain
	middlewares = append([]Middleware{errorHandlerMiddleware}, middlewares...) // 1th
	
	if config.IsDev {
		middlewares = append([]Middleware{loggerMiddleware}, middlewares...)   // 0th
	}

	

	// Execute the middleware in the same order and return the final func.
	// This is a confusing and tricky construct :)
	// We need to use the reverse order since we are chaining inwards.
	for i := len(middlewares) - 1; i >= 0; i-- {
		final = middlewares[i](final) // mw1(mw2(mw3(final)))
	}
	return final
}



func NormalizedMiddleware(handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = handler(w, r) 
	}
}




func errorHandlerMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *errors.Error {
		defer func() {
			if rec := recover(); rec != nil {
				log.Println("Panic Recovered:", rec)
				WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
			}
		}()

		err := next(w, r)
		
		if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
			return nil
		}

		// Custom error handling
		msg := err.Message
		if msg == "" {
			msg = http.StatusText(err.Code)
		}

		if err.Err != nil {
			WriteErrorResponse(w, err.Code, msg, err.Err)
		} else {
			WriteErrorResponse(w, err.Code, msg)
		}

		return err
	}
}






// logger measures request time, status, and response size.
func loggerMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *errors.Error {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}
		err := next(lrw, r)

		duration := time.Since(start)

		// Color codes
		reset := "\033[0m"
		bold := "\033[1m"
		gray := "\033[90m"

		green := "\033[32m"
		yellow := "\033[33m"
		red := "\033[31m"
		cyan := "\033[36m"
		magenta := "\033[35m"

		// Status color based on code
		statusColor := green
		switch {
		case lrw.status >= 500:
			statusColor = red
		case lrw.status >= 400:
			statusColor = yellow
		case lrw.status >= 300:
			statusColor = cyan
		}

		// Build colored log
		log.Printf(
			"%s%-6s%s %s%-40s%s %s%d%s %s%v%s - %s%dB%s",
			bold, r.Method, reset,
			cyan, r.RequestURI, reset,
			statusColor, lrw.status, reset,
			gray, duration, reset,
			magenta, lrw.size, reset,
		)

		return err
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
