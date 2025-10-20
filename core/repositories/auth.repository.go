package repo

import (
	"database/sql"
	"net/http"
	"opsie/config"
	"opsie/core/dbutils"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
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


func (r *AuthRepository) CreateSession(userId types.ID, key string, expiry time.Time) (types.Session, *errors.Error) {
	query := `
		INSERT INTO sessions (user_id, key, expiry)
		VALUES ($1, $2, $3)
		RETURNING ` + dbutils.SessionColumns + `;`

	row := r.db.QueryRow(query, userId, key, expiry)
	return dbutils.SessionScan(row)
}


func (r *AuthRepository) GetValidSessionByKey(key string) (types.Session, *errors.Error) {
	query := `
		SELECT id, user_id, key, ip, os, device, browser, is_valid, expiry, created_at
		FROM sessions
		WHERE key=$1 AND is_valid=true AND expiry > now()
	`

	row := r.db.QueryRow(query, key)
	session, err := dbutils.SessionScan(row)
	if err != nil {
		return types.Session{}, err
	}

	return session, nil
}




func (r *AuthRepository) GetValidSessionWithAuthUser(key string) (types.SessionWithUser, *errors.Error) {
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
        &su.AuthUser.ID,
        &su.AuthUser.DisplayName,
        &su.AuthUser.Email,
        &su.AuthUser.SystemRole,
        &su.AuthUser.IsActive,
    )
    if err != nil {
		if err == sql.ErrNoRows {
			return types.SessionWithUser{}, errors.Unauthorized("invalid session")
		}
        return types.SessionWithUser{}, errors.Internal(err)
    }

    // convert NullStrings
    su.Session.IP = ip.String
    su.Session.OS = os.String
    su.Session.Device = device.String
    su.Session.Browser = browser.String

    if !su.AuthUser.IsActive {
        return types.SessionWithUser{}, errors.New(http.StatusForbidden, "user is inactive")
    }

    return su, nil
}


func (r *AuthRepository) ExpireSession(key string) *errors.Error {
	query := `
		UPDATE sessions
		SET is_valid = false, expiry = NOW()
		WHERE key = $1 AND is_valid = true
	`
	result, err := r.db.Exec(query, key)
	if err != nil {
		return errors.Internal(err)
	}

	// Optional: check if any row was actually updated
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFound("No active session found")
	}

	return nil
}



func (r *AuthRepository) RegenerateSessionKey(key types.SessionKey) (types.Session, *errors.Error) {

	newKey, gskRrr := utils.GenerateSessionKey()
	if gskRrr != nil {
		return types.Session{}, errors.Internal(gskRrr)
	}

	expiry := time.Now().Add(time.Duration(config.AppConfig.SessionDays) * 24 * time.Hour)

	query := `
		UPDATE sessions
		SET key = $2, expiry = $3
		WHERE key = $1 AND is_valid = true
		RETURNING ` + dbutils.SessionColumns + `;
	`

	row := r.db.QueryRow(query, key, newKey, expiry)
	return dbutils.SessionScan(row)
}
