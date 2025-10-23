package db

import (
	"fmt"
	"log"
	"opsie/config"
	"opsie/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// Postgres initializes and returns a GORM DB instance.
// Call this once at startup and reuse the *gorm.DB across your app.
func Postgres() (*gorm.DB, error) {
	env := config.ENV

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		env.PG_Host,
		env.PG_User,
		env.PG_Password,
		env.PG_Database,
		env.PG_Port,
	)

	// Configure GORM logger (optional)
	newLogger := gorm_logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		gorm_logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gorm_logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Connect to Postgres via GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	// Check connection
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("❌ Failed to get generic DB: %v", err)
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		logger.Fatalf("❌ Failed to ping PostgreSQL: %v", err)
		return nil, err
	}

	logger.Info("✅ PostgreSQL connected → [%s:%s/%s]", env.PG_Host, env.PG_Port, env.PG_Database)
	return db, nil
}
