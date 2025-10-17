package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"opsie/constant"
	"opsie/pkg/errors"
	"opsie/pkg/utils"

	"github.com/lib/pq"
)

// Repository - Handles DB operations for user.
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


func (r *Repository) IsUserExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		return false, errors.New(http.StatusInternalServerError, "failed to check user existence")
	}
	return exists, nil
}



func (r *Repository) CreateOwnerAccount(payload TNewOwnerPayload) (TUser, *errors.Error) {
	query := `
		INSERT INTO users (id, display_name, email, password, system_role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, display_name, email, system_role, preference, is_active, created_at, updated_at;
	`

	var user TUser
	var prefBytes []byte

	id := utils.GenerateID()
	system_role := constant.SystemRoleOwner

	err := r.db.QueryRow(query, id, payload.DisplayName, payload.Email, payload.Password, system_role).Scan(
		&user.Id,
		&user.DisplayName,
		&user.Email,
		&user.SystemRole,
		&prefBytes,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		// Detect duplicate email constraint
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return TUser{}, errors.New(409, "email already in use")
		}
		
		return TUser{}, errors.Internal(err)
	}

	// Decode preference JSONB
	if len(prefBytes) > 0 {
		_ = json.Unmarshal(prefBytes, &user.Preference)
	} else {
		user.Preference = make(map[string]any)
	}

	return user, nil
}

