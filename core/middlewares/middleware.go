package mw

import (
	"database/sql"
	"opsie/core/repo"
	"opsie/pkg/bolt"
)

var Auth bolt.Middleware

func Register(db *sql.DB)  {
	// Repositories
	authRepo := repo.NewAuthRepository(db)
	
	// Middlewares
	Auth = newAuthMiddleware(authRepo)

}

