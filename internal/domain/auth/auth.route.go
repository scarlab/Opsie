package auth

import (
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
//
// Example:
//   r.HandleFunc("/get/items", h.GetItems).Methods("GET")
func HandleRoutes(r *mux.Router, h *Handler) {
	// Example:
	// r.HandleFunc("/get/something", h.GetSomething).Methods("GET")
	
	
	bolt.Api(r, bolt.MethodPost, "/login", h.Login)
	bolt.Api(r, bolt.MethodPost, "/create", h.Logout)
}
