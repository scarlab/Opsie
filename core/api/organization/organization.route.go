package organization

import (
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r *mux.Router, h *Handler) {
	bolt.Api(r, bolt.MethodPost, "/create", h.Create)
}

