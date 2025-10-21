package auth

import (
	mw "opsie/core/middlewares"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r chi.Router, h *Handler) {
	api.Post(r, "/login", h.Login)
	
	// MW - [Auth]
	api.Get(r, "/session", h.GetSessionUser, mw.Auth)
	api.Get(r, "/logout", h.Logout, mw.Auth)
}
