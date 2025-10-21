package mw

import (
	"net/http"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/logger"
	"opsie/types"
	"reflect"
	"runtime"
)



func Recoverer(next types.HandlerFunc) types.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *errors.Error {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("Panic: %s", rec)
				bolt.WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
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
			bolt.WriteErrorResponse(w, err.Code, msg, err.Err)
		} else {
			bolt.WriteErrorResponse(w, err.Code, msg)
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
