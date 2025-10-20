package dbutils

import (
	"database/sql"
	"encoding/json"
	"opsie/pkg/errors"
	"opsie/types"
)

const UserColumns = `
    id,
    display_name,
    email,
    password,
    avatar,
    system_role,
    preference,
    is_active,
    updated_at,
    created_at
`

func UserScan(row *sql.Row) (types.User, *errors.Error) {
	var user types.User
	var prefBytes []byte
	var avatar sql.NullString

	err := row.Scan(
		&user.ID,
		&user.DisplayName,
		&user.Email,
		&user.Password,
		&avatar,
		&user.SystemRole,
		&prefBytes,
		&user.IsActive,
		&user.UpdatedAt,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, errors.NotFound("User not found")
		}
		return types.User{}, errors.Internal(err)
	}

	// Handle nullable avatar
	user.Avatar = ""
	if avatar.Valid {
		user.Avatar = avatar.String
	}

	// Unmarshal preferences JSON
	if len(prefBytes) > 0 {
		_ = json.Unmarshal(prefBytes, &user.Preference)
	} else {
		user.Preference = make(map[string]any)
	}

	return user, nil
}