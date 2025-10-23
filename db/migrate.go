package db

import (
	"opsie/core/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Team{},
		&models.UserTeam{},
		&models.Node{},
		&models.Project{},
		&models.Resource{},
		&models.ResourceNode{},
	)
}

