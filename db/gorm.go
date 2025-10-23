package db

import (
	"opsie/core/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Team{},
		&models.UserTeam{},
		&models.Node{},
		&models.Project{},
		&models.Resource{},
		&models.ResourceNode{},
	)
	return db, err
}
