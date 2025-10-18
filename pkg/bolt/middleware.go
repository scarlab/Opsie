package bolt

import (
	"fmt"
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
		fmt.Println("err",err)
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
	return func(w http.ResponseWriter, r *http.Request) *errors.Error{
		start := time.Now()

		// Wrap the ResponseWriter
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

		next(lrw, r) // call the next handler

		duration := time.Since(start)
		log.Printf("%s %s %d %v %d", r.Method, r.RequestURI, lrw.status, duration, lrw.size)
		errors.New(http.StatusForbidden, "user is inactiv exx")
		return nil
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
