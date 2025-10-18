package mw

import (
	"fmt"
	"net/http"
	repo "opsie/core/repositories"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"strings"
	"time"
)


func newAuthMiddleware(authRepo *repo.AuthRepository) bolt.Middleware {
	return func(next bolt.HandlerFunc) bolt.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) *errors.Error {
			// Get the session key from cookie/header
			// 1. Try cookie
			var sessionKey string
			if cookie, err := r.Cookie("session"); err == nil {
				sessionKey = cookie.Value
			}

			// 2. Fallback to header
			if sessionKey == "" {
				sessionKey = r.Header.Get("X-Session-Key")
				if sessionKey == "" {
					authHeader := r.Header.Get("Authorization")
					if strings.HasPrefix(authHeader, "Bearer ") {
						sessionKey = strings.TrimPrefix(authHeader, "Bearer ")
					}
				}
			}

			// 3. Validate
			if sessionKey == "" {
				return errors.New(http.StatusUnauthorized, "missing session key")
			}

			fmt.Println("{auth} session key:", sessionKey)

			authRepo.CreateSession(2, "", time.Now())

			// TODO: Fetch & validate session in DB

			next(w, r)
			return nil
		}
	}
}