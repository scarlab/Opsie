package {{.PackageName}}

import "github.com/gorilla/mux"

// RegisterRoutes - Defines all HTTP endpoints for {{.PackageName}}.
//
// Example:
//   r.HandleFunc("/get/items", h.GetItems).Methods("GET")
func RegisterRoutes(r *mux.Router, h *Handler) {
	// Example:
	// r.HandleFunc("/get/something", h.GetSomething).Methods("GET")
	
	r.HandleFunc("/health", h.Health).Methods("GET") // Health checkup route
}
