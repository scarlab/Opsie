package repo

import (
	"database/sql"
	"opsie/core/dbutils"
	"opsie/def"
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
	query := `INSERT INTO users (id, display_name, email, password, system_role)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING ` + dbutils.UserColumns 

	var user types.User

	id := utils.GenerateID()
	system_role := def.SystemRoleOwner

	row := r.db.QueryRow(query, id, payload.DisplayName, payload.Email, payload.Password, system_role)
	 
	user, err := dbutils.UserScan(row)
	if err != nil {
		if pqErr, ok := err.Original().(*pq.Error); ok && pqErr.Code == "23505" {
			return types.User{}, errors.New(409, "Email already in use")
		}
		return types.User{}, err
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
	query := `SELECT ` + dbutils.UserColumns + ` FROM users WHERE email = $1`
	
	row := r.db.QueryRow(query, email)
	return dbutils.UserScan(row)
}




func (r *UserRepository) GetByID(ID types.ID) (types.User, *errors.Error) {
	query := `SELECT ` + dbutils.UserColumns + ` FROM users WHERE id = $1`

	row := r.db.QueryRow(query, ID)
	return dbutils.UserScan(row)
}





func (r *UserRepository) UpdateAccountName(userID types.ID, name string) (types.User, *errors.Error) {
	query := `UPDATE users SET display_name = $1 WHERE id = $2 RETURNING ` + dbutils.UserColumns 


	row := r.db.QueryRow(query, name, userID)
	return dbutils.UserScan(row)
}




func (r *UserRepository) UpdateAccountPassword(userID types.ID, password string) (bool, *errors.Error) {
	query := ` UPDATE users SET password = $1 WHERE id = $2`

	res, err := r.db.Exec(query, password, userID)
	if err != nil {
		return false, errors.Internal(err)
	}

	rowsAffected, _ := res.RowsAffected()
	return rowsAffected > 0, nil
}



