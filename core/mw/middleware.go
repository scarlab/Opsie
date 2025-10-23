package mw

import (
	"opsie/core/repo"
	"opsie/types"

	"gorm.io/gorm"
)

var Auth types.Middleware

func Register(db *gorm.DB)  {
	// Repositories
	authRepo := repo.NewAuthRepository(db)
	
	// Middlewares
	Auth = newAuthMiddleware(authRepo)

}

