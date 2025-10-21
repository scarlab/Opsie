package mw

import (
	"database/sql"
	"opsie/core/repo"
	"opsie/types"
)

var Auth types.Middleware

func Register(db *sql.DB)  {
	// Repositories
	authRepo := repo.NewAuthRepository(db)
	
	// Middlewares
	Auth = newAuthMiddleware(authRepo)

}

