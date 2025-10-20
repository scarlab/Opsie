package dbutils

import (
	"database/sql"
	"opsie/pkg/errors"
	"opsie/types"
)

// SessionColumns defines the consistent list of columns to select or return.
const SessionColumns = `
	id,
	user_id,
	key,
	ip,
	os,
	device,
	browser,
	is_valid,
	expiry,
	created_at
`

// SessionScan scans a single sql.Row into a Session struct.
func SessionScan(row *sql.Row) (types.Session, *errors.Error) {
	var s types.Session
	var ip, os, device, browser sql.NullString

	err := row.Scan(
		&s.ID,
		&s.UserID,
		&s.Key,
		&ip,
		&os,
		&device,
		&browser,
		&s.IsValid,
		&s.Expiry,
		&s.CreatedAt,
	)
	if err != nil {
		// handle known SQL errors here, if you want centralization
		if err == sql.ErrNoRows {
			return types.Session{}, errors.NotFound("Session not found")
		}
		return types.Session{}, errors.Internal(err)
	}

	s.IP = ip.String
	s.OS = os.String
	s.Device = device.String
	s.Browser = browser.String

	return s, nil
}


