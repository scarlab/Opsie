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
	bolt.Api(r, bolt.MethodPost, "/owner/create", h.CreateOwnerAccount)
	bolt.Api(r, bolt.MethodGet, "/owner/count", h.GetOwnerCount)

	// Auth

}
