package mw

import (
	"context"
	"net/http"
	"opsie/constant"
	repo "opsie/core/repositories"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"strings"
)

func newAuthMiddleware(authRepo *repo.AuthRepository) bolt.Middleware {
	return func(next bolt.HandlerFunc) bolt.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) *errors.Error {
			// 1. Get session key
			var sessionKey string
			if cookie, err := r.Cookie("session"); err == nil {
				sessionKey = cookie.Value
			}
			if sessionKey == "" {
				sessionKey = r.Header.Get("X-Session-Key")
				if sessionKey == "" {
					authHeader := r.Header.Get("Authorization")
					if strings.HasPrefix(authHeader, "Bearer ") {
						sessionKey = strings.TrimPrefix(authHeader, "Bearer ")
					}
				}
			}

			if sessionKey == "" {
				return errors.New(http.StatusUnauthorized, "missing session key")
			}

			// 2. Fetch session + user in a single query
			sessionUser, err := authRepo.GetValidSessionWithUser(sessionKey)
			if err != nil {
				return err
			}
			if !sessionUser.User.IsActive {
				return errors.New(http.StatusForbidden, "user is inactive")
			}

			// 3. Attach to context
			ctx := context.WithValue(r.Context(), constant.ContextKeySession, sessionUser.Session)
			ctx = context.WithValue(ctx, constant.ContextKeyUser, sessionUser.User)
			r = r.WithContext(ctx)

			// 4. Call next middleware/handler
			return next(w, r)
		}
	}
}
