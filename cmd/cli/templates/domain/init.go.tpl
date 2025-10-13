// {{.CreatedAt}}

package {{.PackageName}}

import (
	"database/sql"
	"opsie/internal/socket"

	"github.com/gorilla/mux"
)

// Init - Entry point for initializing domain - {{.PackageName}}
//
// Responsibilities:
// 1. Create repository, service, and handler instances.
// 2. Wire dependencies in the correct order.
// 3. Register domain-specific routes.
//
// Usage:
//   packagename.Register(router, db)
func Register(r *mux.Router, db *sql.DB, socketHub *socket.Hub) {
	// Step 1: Create repository (DB layer)
	repository := NewRepository(db)

	// Step 2: Create service (Business logic layer)
	service := NewService(repository)

	// Step 3: Create handler (HTTP layer)
	handler := NewHandler(service, socketHub)

	// Step 4: Create the sub-router for this domain (modify if required)
	router := r.PathPrefix("/{{.PackageName}}").Subrouter()

	// Step 5: Register routes for this domain
	RegisterRoutes(router, handler)
}
