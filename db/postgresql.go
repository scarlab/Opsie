package db

import (
	"database/sql"
	"fmt"
	"log"
	"watchtower/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func Postgres() (*sql.DB, error) {
	// Build DSN (connection string)
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.ENV.PG_USER,
		config.ENV.PG_PASSWD,
		config.ENV.PG_HOST,
		config.ENV.PG_PORT,
		config.ENV.PG_DB,
	)

	// Connect
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open postgres connection:", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping postgres:", err)
	}

	log.Printf("✔️  PostgreSQL Connected [%s]", config.ENV.PG_DB)
	return db, nil
}
