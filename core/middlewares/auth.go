package mw

import (
	"context"
	"net/http"
	repo "opsie/core/repositories"
	"opsie/def"
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
				return errors.New(http.StatusUnauthorized, "unauthorized")
			}

			// 2. Fetch session + user in a single query
			authUser, err := authRepo.GetValidSessionWithAuthUser(sessionKey)
			if err != nil {
				return err
			}


			// 3. Attach to context
			ctx := context.WithValue(r.Context(), def.ContextKeySession, authUser.Session)
			ctx = context.WithValue(ctx, def.ContextKeyUser, authUser.AuthUser)
			r = r.WithContext(ctx)

			// 4. Call next middleware/handler
			return next(w, r)
		}
	}
}
