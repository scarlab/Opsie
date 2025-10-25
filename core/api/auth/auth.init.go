// 2025/10/17 16:59:02

package auth

import (
	"opsie/core/repo"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// Init - Entry point for initializing api - auth
//
// Responsibilities:
// 1. Create repository(s), handler instances.
// 2. Wire dependencies in the correct order.
// 3. Register api-specific routes.

func Register(r chi.Router, db *gorm.DB) {
	// Step 1: Create repository (DB layer)
	repository := repo.NewAuthRepository(db)
	userRepository := repo.NewUserRepository(db)

	// Step 2: Create handler (HTTP layer)
	handler := NewHandler(repository, userRepository)

	// Step 3: Register routes for this api
	HandleRoutes(r, handler) 
}
