package user

import "github.com/gorilla/mux"

// RegisterRoutes - Defines all HTTP endpoints for user.
func RegisterRoutes(r *mux.Router, h *Handler) {

	// Public --------------------------------------------------
	r.HandleFunc("/health", h.Health).Methods("GET")
	r.HandleFunc("/create/owner", h.CreateOwnerAccount).Methods("POST")


}
