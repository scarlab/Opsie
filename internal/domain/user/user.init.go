// 2025/10/14 12:59:19

package user

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Init - Entry point for initializing domain - user
//
// Responsibilities:
// 1. Create repository, service, and handler instances.
// 2. Wire dependencies in the correct order.
// 3. Register domain-specific routes.
//
// Usage:
//   packagename.Register(router, db)
func Register(r *mux.Router, db *sql.DB) {
	// Step 1: Create repository (DB layer)
	repository := NewRepository(db)

	// Step 2: Create service (Business logic layer)
	service := NewService(repository)

	// Step 3: Create handler (HTTP layer)
	handler := NewHandler(service)

	// Step 4: Create the sub-router for this domain (modify if required)
	router := r.PathPrefix("/user").Subrouter()

	// Step 5: Register routes for this domain
	HandleRoutes(router, handler)
}
