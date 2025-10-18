package mw

import (
	"net/http"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
)


func Auth(next bolt.THandlerFunc) bolt.THandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *errors.Error {
		//TODO:
		// Get Session Key `session` from cookie & header
		// Fetch the session data by `sessions.key`
		// Validate
		

		
		next(w, r)
		return nil
	}
}

