package mw

import (
	"net/http"

	"opsie/def"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/types"
)

func Admin(next types.HandlerFunc) types.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) *errors.Error {
			// Get the request/session user 
			sessionUser, gsuErr := bolt.GetSessionUser(r)
			if gsuErr!= nil {return gsuErr}

			// Check if the user has admin privilege or not
			if sessionUser.SystemRole == def.SystemRoleStaff.ToString() {
				return errors.Forbidden("I think you're not supposed to be here...")
			}

			// 4. Call next middleware/handler
			return next(w, r)
		}
	}
