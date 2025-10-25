// 2025/10/14 12:59:19

package user

import (
	"opsie/core/repo"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
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
func Register(r chi.Router, db *gorm.DB) {
	// Step 1: Create repository (DB layer)
	repository := repo.NewUserRepository(db)
	authRepository := repo.NewAuthRepository(db)
	teamRepository := repo.NewTeamRepository(db)
	userTeamRepository := repo.NewUserTeamRepository(db)

	// Step 2: Create handler (HTTP layer)
	handler := NewHandler(repository, authRepository, teamRepository, userTeamRepository)

	// Step 3: Register routes for this api
	HandleRoutes(r, handler) 
}
