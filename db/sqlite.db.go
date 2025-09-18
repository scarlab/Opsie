package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"watchtower/config"

	_ "modernc.org/sqlite" // SQLite driver
)

// SQLite opens (or creates) a SQLite database file.
func SQLite() (*sql.DB, error) {

	path := config.ENV.MainDBPath

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, fmt.Errorf("failed to create db directory: %w", err)
	}
	
	// Open or create file
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite db: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping sqlite db: %w", err)
	}

	return db, nil
}
