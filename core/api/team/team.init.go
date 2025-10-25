// 2025/10/20 15:33:22

package team

import (
	"opsie/core/repo"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// Init - Entry point for initializing api - team
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
	repository := repo.NewTeamRepository(db)
	userTeamRepository := repo.NewUserTeamRepository(db)

	// Step 2: Create handler (HTTP layer)
	handler := NewHandler(repository, userTeamRepository)

	// Step 3: Register routes for this api
	HandleRoutes(r, handler) 
}
