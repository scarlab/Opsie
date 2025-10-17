package auth

import (
	"database/sql"
	"opsie/pkg/errors"
	"time"
)

// Repository - Handles DB operations for auth.
// Talks only to the database (or other data sources).
type Repository struct {
	db *sql.DB
}

// NewRepository - Constructor for Repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}


func (r *Repository) CreateSession(userId int64, key string, expiry time.Time) (TSession, *errors.Error) {
	query := `
		INSERT INTO sessions (user_id, key, expiry)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, key, ip, os, device, browser, is_valid, expiry, created_at;
	`

	var session TSession
	var ip, os, device, browser sql.NullString

	err := r.db.QueryRow(query, userId, key, expiry).Scan(
		&session.ID,
		&session.UserID,
		&session.Key,
		&ip,
		&os,
		&device,
		&browser,
		&session.IsValid,
		&session.Expiry,
		&session.CreatedAt,
	)
	if err != nil {
		return TSession{}, errors.Internal(err)
	}

	// Convert NullString to normal string
	session.IP = ip.String
	session.OS = os.String
	session.Device = device.String
	session.Browser = browser.String


	return session, nil
}
