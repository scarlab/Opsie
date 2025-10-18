// 2025/10/17 16:59:02

package auth

import (
	"database/sql"
	repo "opsie/core/repositories"
	"opsie/core/services"

	"github.com/gorilla/mux"
)

// Init - Entry point for initializing domain - auth
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
	repository := repo.NewAuthRepository(db)
	userRepository := repo.NewUserRepository(db)

	// Step 2: Create service (Business logic layer)
	service := services.NewAuthService(repository, userRepository)

	// Step 3: Create handler (HTTP layer)
	handler := NewHandler(service)

	// Step 4: Create the sub-router for this domain (modify if required)
	router := r.PathPrefix("/auth").Subrouter()

	// Step 5: Register routes for this domain
	HandleRoutes(router, handler)
}
