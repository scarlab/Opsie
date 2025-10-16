package user

import (
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for user.
func HandleRoutes(r *mux.Router, h *Handler) {

	// Public --------------------------------------------------
	// r.HandleFunc("/health", h.Health).Methods("GET")
	// r.HandleFunc("/create/owner", h.CreateOwnerAccount).Methods("POST")
	bolt.Api(r, bolt.GET, "/health", h.Health)
	bolt.Api(r, bolt.POST, "/create/owner", h.CreateOwnerAccount)

}
