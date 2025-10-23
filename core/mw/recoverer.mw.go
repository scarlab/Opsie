package mw

import (
	"net/http"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/logger"
	"opsie/types"
	"reflect"
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
			bolt.WriteErrorResponse(w, err.Code, msg, err.Err)
		} else {
			bolt.WriteErrorResponse(w, err.Code, msg)
		}

		return err
	}
}


// // helper to get caller file:line
// func getCaller(s int) (string, int) {
//     _, file, line, ok := runtime.Caller(s) // 2 to skip runtime + logger function
//     if !ok {
//         return "unknown", 0
//     }
//     return file, line
// }

// // getErrorOrigin returns the file and line of the first caller outside runtime/logging packages
// func getErrorOrigin() (file string, line int) {
//     // skip 0 = this function, 1 = caller, 2 = caller's caller...
//     for i := 1; i < 20; i++ {
//         pc, f, l, ok := runtime.Caller(i)
//         if !ok {
//             break
//         }
//         funcName := runtime.FuncForPC(pc).Name()
//         // skip standard library / logging / error helper frames
//         if !strings.Contains(funcName, "runtime.") &&
//             !strings.Contains(funcName, "log") &&
//             !strings.Contains(funcName, "errors") {
//             return f, l
//         }
//     }
//     return "unknown", 0
// }
