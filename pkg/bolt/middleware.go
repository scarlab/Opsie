package bolt

import (
	"net/http"
	"opsie/pkg/errors"
	"opsie/pkg/logger"
	"reflect"
	"runtime"
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
	
	// Add Error Handler & Logger to the chain
	middlewares = append([]Middleware{loggerMiddleware, errorHandlerMiddleware}, middlewares...) // 1th
		

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
				logger.Error("Panic: %s", rec)
				WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
			}
		}()

		err := next(w, r)
		
		if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
			return nil
		}

		// Custom error handling
		msg := err.Error
		if msg == "" {
			msg = http.StatusText(err.Code)
		}


		if err.Err != nil {
			file, line := getCaller()

			if err.Code ==500 {
				logger.Error("%s", err.Err)
				logger.Printf("%s at %d", file, line)
			}
			WriteErrorResponse(w, err.Code, msg, err.Err)
		} else {
			WriteErrorResponse(w, err.Code, msg)
		}

		return err
	}
}


// helper to get caller file:line
func getCaller() (string, int) {
    _, file, line, ok := runtime.Caller(2) // 2 to skip runtime + logger function
    if !ok {
        return "unknown", 0
    }
    return file, line
}




// logger measures request time, status, and response size.
func loggerMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *errors.Error {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

		err := next(lrw, r)
		duration := time.Since(start)

		statusColor := colorForStatus(lrw.status)
		methodColor := colorForMethod(r.Method)
		sizeColor := "\033[36m"
		reset := "\033[0m"
		bold := "\033[1m"

		fLogger, sLogger := logger.HttpLogger()

		sLogger.Printf("%s%s%-6s%s  %-40s  %s%3d%s  %8.2fms  %s%dB%s\n",
			bold, methodColor, r.Method, reset,
			r.URL.Path,
			statusColor, lrw.status, reset,
			float64(duration.Microseconds())/1000.0,
			sizeColor, lrw.size, reset,
		)

		fLogger.Printf("%-6s  %-40s  %3d  %8.2fms  %dB\n",
			r.Method,
			r.URL.Path,
			lrw.status,
			float64(duration.Microseconds())/1000.0,
			lrw.size,
		)

		return err
	}
}

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
	n, err := lrw.ResponseWriter.Write(b)
	lrw.size += n
	return n, err
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[32m" // green
	case code >= 300 && code < 400:
		return "\033[36m" // cyan
	case code >= 400 && code < 500:
		return "\033[33m" // yellow
	default:
		return "\033[31m" // red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return "\033[34m" // blue
	case "POST":
		return "\033[32m" // green
	case "PUT":
		return "\033[33m" // yellow
	case "PATCH":
		return "\033[33m" // yellow
	case "DELETE":
		return "\033[31m" // red
	default:
		return "\033[37m" // white
	}
}