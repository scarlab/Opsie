package user

import (
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for user.
func HandleRoutes(r *mux.Router, h *Handler) {

	// ---------------------------------------------------------
	// PUBLIC...
	// ---------------------------------------------------------
	// Onboarding
	bolt.Api(r, bolt.POST, "/create/owner", h.CreateOwnerAccount)

	// Auth

}
