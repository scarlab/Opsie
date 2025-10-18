package auth

import (
	mw "opsie/core/middlewares"
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r *mux.Router, h *Handler) {
	bolt.Api(r, bolt.MethodPost, "/login", h.Login)
	bolt.Api(r, bolt.MethodGet, "/logout", h.Logout, mw.Auth)
}
