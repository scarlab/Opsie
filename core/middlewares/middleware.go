package mw

import (
	"database/sql"
	repo "opsie/core/repositories"
	"opsie/pkg/bolt"
)

var Auth bolt.Middleware

func Register(db *sql.DB)  {
	// Repositories
	authRepo := repo.NewAuthRepository(db)
	
	// Middlewares
	Auth = newAuthMiddleware(authRepo)

}

