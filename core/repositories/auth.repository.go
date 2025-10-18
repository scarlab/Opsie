package repo

import (
	"database/sql"
	"net/http"
	"opsie/pkg/errors"
	"opsie/types"
	"time"
)

// AuthRepository - Handles DB operations for auth.
// Talks only to the database (or other data sources).
type AuthRepository struct {
	db *sql.DB
}

// NewAuthRepository - Constructor for Repository
func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}


func (r *AuthRepository) CreateSession(userId int64, key string, expiry time.Time) (types.Session, *errors.Error) {
	query := `
		INSERT INTO sessions (user_id, key, expiry)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, key, ip, os, device, browser, is_valid, expiry, created_at;
	`

	var session types.Session
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
		return types.Session{}, errors.Internal(err)
	}

	// Convert NullString to normal string
	session.IP = ip.String
	session.OS = os.String
	session.Device = device.String
	session.Browser = browser.String


	return session, nil
}

func (r *AuthRepository) GetValidSessionByKey(key string) (types.Session, *errors.Error) {
	query := `
		SELECT id, user_id, key, ip, os, device, browser, is_valid, expiry, created_at
		FROM sessions
		WHERE key=$1 AND is_valid=true AND expiry > now()
	`

	var session types.Session
	var ip, os, device, browser sql.NullString

	err := r.db.QueryRow(query, key).Scan(
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
		if err == sql.ErrNoRows {
			return types.Session{}, errors.NotFound("session not found or expired")
		}
		return types.Session{}, errors.Internal(err)
	}

	// Convert NullString to normal string
	session.IP = ip.String
	session.OS = os.String
	session.Device = device.String
	session.Browser = browser.String

	return session, nil
}




func (r *AuthRepository) GetValidSessionWithUser(key string) (types.SessionWithUser, *errors.Error) {
    var su types.SessionWithUser
    var ip, os, device, browser sql.NullString

    query := `
        SELECT s.id, s.user_id, s.key, s.ip, s.os, s.device, s.browser, s.is_valid, s.expiry, s.created_at,
               u.id, u.display_name, u.email, u.system_role, u.is_active
        FROM sessions s
        JOIN users u ON u.id = s.user_id
        WHERE s.key = $1 AND s.is_valid = true AND s.expiry > NOW() AND u.is_active = true
    `

    err := r.db.QueryRow(query, key).Scan(
        &su.Session.ID,
        &su.Session.UserID,
        &su.Session.Key,
        &ip, &os, &device, &browser,
        &su.Session.IsValid,
        &su.Session.Expiry,
        &su.Session.CreatedAt,
        &su.User.ID,
        &su.User.DisplayName,
        &su.User.Email,
        &su.User.SystemRole,
        &su.User.IsActive,
    )
    if err != nil {
        return types.SessionWithUser{}, errors.Internal(err)
    }

    // convert NullStrings
    su.Session.IP = ip.String
    su.Session.OS = os.String
    su.Session.Device = device.String
    su.Session.Browser = browser.String

    if !su.User.IsActive {
        return types.SessionWithUser{}, errors.New(http.StatusForbidden, "user is inactive")
    }

    return su, nil
}
