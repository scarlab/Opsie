// {{.CreatedAt}}

package {{.PackageName}}

import (
	"database/sql"
	"opsie/core/socket"
	"opsie/core/repo"
	"opsie/core/services"

	"github.com/go-chi/chi/v5"
)

// Init - Entry point for initializing api - {{.PackageName}} with socket hub
//
// Responsibilities:
// 1. Create repository, service, and handler instances.
// 2. Wire dependencies in the correct order.
// 3. Register api-specific routes.
//
// Usage:
//   packagename.Register(router, db)
func Register(r chi.Router, db *sql.DB, socketHub *socket.Hub) {
	// Step 1: Create repository (DB layer)
	repository := repo.New{{.Name}}Repository(db)

	// Step 2: Create service (Business logic layer)
	service := services.New{{.Name}}Service(repository)

	// Step 3: Create handler (HTTP layer)
	handler := NewHandler(service, socketHub)

	// Step 4: Register routes for this api
	HandleRoutes(r, handler)
}
