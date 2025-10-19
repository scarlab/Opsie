package repo

import (
	"database/sql"
	"encoding/json"
	"opsie/constant"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"

	"github.com/lib/pq"
)

// UserRepository - Handles DB operations for user.
// Talks only to the database (or other data sources).
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository - Constructor for Repository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}



func (r *UserRepository) CreateOwnerAccount(payload types.NewOwnerPayload) (types.User, *errors.Error) {
	query := `
		INSERT INTO users (id, display_name, email, password, system_role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, display_name, email, system_role, preference, is_active, created_at, updated_at;
	`

	var user types.User
	var prefBytes []byte

	id := utils.GenerateID()
	system_role := constant.SystemRoleOwner

	err := r.db.QueryRow(query, id, payload.DisplayName, payload.Email, payload.Password, system_role).Scan(
		&user.ID,
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
			return types.User{}, errors.New(409, "email already in use")
		}
		
		return types.User{}, errors.Internal(err)
	}

	// Decode preference JSONB
	if len(prefBytes) > 0 {
		_ = json.Unmarshal(prefBytes, &user.Preference)
	} else {
		user.Preference = make(map[string]any)
	}

	return user, nil
}

func (r *UserRepository) GetOwnerCount() (int, *errors.Error) {
    var count int

    query := `SELECT COUNT(*) FROM users WHERE system_role = 'owner'`

    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, errors.Internal(err)
    }

    return count, nil
}




func (r *UserRepository) GetByEmail(email string) (types.User, *errors.Error) {
	var user types.User
	query := `
		SELECT id, display_name, email, password, system_role, preference, is_active, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	var prefBytes []byte

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.DisplayName,
		&user.Email,
		&user.Password,
		&user.SystemRole,
		&prefBytes,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, errors.NotFound("user not found")
		}
		return types.User{}, errors.Internal(err)
	}

	if len(prefBytes) > 0 {
		_ = json.Unmarshal(prefBytes, &user.Preference)
	} else {
		user.Preference = make(map[string]any)
	}

	return user, nil
}




func (r *UserRepository) GetByID(ID int64) (types.User, *errors.Error) {
	var user types.User
	query := `
		SELECT id, display_name, email, password, system_role, preference, is_active, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	var prefBytes []byte

	err := r.db.QueryRow(query, ID).Scan(
		&user.ID,
		&user.DisplayName,
		&user.Email,
		&user.Password,
		&user.SystemRole,
		&prefBytes,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, errors.NotFound("user not found")
		}
		return types.User{}, errors.Internal(err)
	}

	if len(prefBytes) > 0 {
		_ = json.Unmarshal(prefBytes, &user.Preference)
	} else {
		user.Preference = make(map[string]any)
	}

	return user, nil
}


