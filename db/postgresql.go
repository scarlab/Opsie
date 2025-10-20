package db

import (
	"database/sql"
	"fmt"
	"opsie/config"
	"opsie/pkg/logger"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Postgres initializes a PostgreSQL connection using environment variables
// defined in config.ENV. It logs connection progress and returns the DB handle.
//
// Note: this function should be called once at startup and the returned
// *sql.DB should be shared throughout the application.
func Postgres() (*sql.DB, error) {
	// log.Println("🔌 Connecting to PostgreSQL...")

	// Build DSN (connection string)
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.ENV.PG_User,
		config.ENV.PG_Password,
		config.ENV.PG_Host,
		config.ENV.PG_Port,
		config.ENV.PG_Database,
	)

	// Establish connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatalf("Failed to open PostgreSQL connection: %v", err)
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		logger.Fatalf("Failed to ping PostgreSQL: %v", err)
		return nil, err
	}

	logger.Info("✅ PostgreSQL connected → [%s:%s/%s]\n", config.ENV.PG_Host, config.ENV.PG_Port, config.ENV.PG_Database)

	return db, nil
}
