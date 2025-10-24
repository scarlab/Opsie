package auth

import (
	"opsie/core/mw"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r chi.Router, h *Handler) {
	
	// ____________________________________________________________________________
	// Public Routes --------------------------------------------------------------
	api.Post(r, 		"/login", 			h.Login)
	
	/// ___________________________________________________________________________
	/// Protected [Auth] ----------------------------------------------------------
	/// User Account Routes
	api.Get(r, 			"/session", 		h.GetSessionUser, 		mw.Auth)
	api.Get(r, 			"/logout", 			h.Logout, 				mw.Auth)
}
