// 2025/10/14 12:59:19

package user

import (
	"database/sql"
	"opsie/core/repo"
	"opsie/core/services"

	"github.com/go-chi/chi/v5"
)

// Init - Entry point for initializing api - user
//
// Responsibilities:
// 1. Create repository, service, and handler instances.
// 2. Wire dependencies in the correct order.
// 3. Register api-specific routes.
//
// Usage:
//   packagename.Register(router, db)
func Register(r chi.Router, db *sql.DB) {
	// Step 1: Create repository (DB layer)
	repository := repo.NewUserRepository(db)
	authRepository := repo.NewAuthRepository(db)
	teamRepository := repo.NewTeamRepository(db)
	userTeamRepository := repo.NewUserTeamRepository(db)

	// Step 2: Create service (Business logic layer)
	service := services.NewUserService(repository, authRepository, teamRepository, userTeamRepository)

	// Step 3: Create handler (HTTP layer)
	handler := NewHandler(service)

	// Step 4: Create the sub-router for this api (modify if required)
	// [v0.0.1-beta] legacy mux router implementation
	// router := r.PathPrefix("/team").Subrouter()

	// Step 5: Register routes for this api
	HandleRoutes(r, handler) // [chi-v0.0.2] new chi router implementation
}
