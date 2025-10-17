package middlewares

import (
	"net/http"
	"opsie/pkg/bolt"
)

func ErrorMiddleware(next  bolt.THandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if next == nil {
			return
		}
		
		next(w, r)
	

		var status int
		var msg string

		// switch {
		// case errors.Is(err, sql.ErrNoRows):
		// 	status = http.StatusNotFound
		// 	msg = "resource not found"
		// case errors.Is(err, ErrUnauthorized):
		// 	status = http.StatusUnauthorized
		// 	msg = "unauthorized"
		// case errors.Is(err, ErrBadRequest):
		// 	status = http.StatusBadRequest
		// 	msg = "bad request"
		// default:
		// 	status = http.StatusInternalServerError
		// 	msg = "internal server error"
		// }

		// log.Printf("ERROR %d: %v", status, err)
		http.Error(w, msg, status)
	}
}
