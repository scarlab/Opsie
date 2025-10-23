package repo

import (
	"context"
	"database/sql"
	"opsie/core/dbutils"
	"opsie/def"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// UserRepository - Handles DB operations for user.
// Talks only to the database (or other data sources).
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository - Constructor for Repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}


func (r *UserRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, *errors.Error) {
    tx, err := r.db.BeginTx(ctx, opts)
    if err != nil {
        return nil, errors.Internal(err)
    }
    return tx, nil
}



func (r *UserRepository) CreateOwnerAccount(tx *sql.Tx, payload types.NewOwnerPayload) (types.UserModel, *errors.Error) {
	query := `
		INSERT INTO users (id, display_name, email, password, system_role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING ` + dbutils.UserColumns

	id := utils.GenerateID()
	system_role := def.SystemRoleOwner

	// Choose the right query executor
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRow(query, id, payload.DisplayName, payload.Email, payload.Password, system_role)
	} else {
		row = r.db.QueryRow(query, id, payload.DisplayName, payload.Email, payload.Password, system_role)
	}

	user, err := dbutils.UserScan(row)
	if err != nil {
		// Handle unique constraint violation
		if pqErr, ok := err.Original().(*pq.Error); ok && pqErr.Code == "23505" {
			return types.UserModel{}, errors.New(409, "Email already in use")
		}
		return types.UserModel{}, err
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


func (r *UserRepository) Delete(userID int64) *errors.Error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, userID)
	if err != nil {
		return errors.Internal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFound("user not found")
	}

	return nil
}




func (r *UserRepository) GetByEmail(email string) (types.UserModel, *errors.Error) {
	query := `SELECT ` + dbutils.UserColumns + ` FROM users WHERE email = $1`
	
	row := r.db.QueryRow(query, email)
	return dbutils.UserScan(row)
}




func (r *UserRepository) GetByID(ID int64) (types.UserModel, *errors.Error) {
	query := `SELECT ` + dbutils.UserColumns + ` FROM users WHERE id = $1`

	row := r.db.QueryRow(query, ID)
	return dbutils.UserScan(row)
}





func (r *UserRepository) UpdateAccountName(userID int64, name string) (types.UserModel, *errors.Error) {
	query := `UPDATE users SET display_name = $1 WHERE id = $2 RETURNING ` + dbutils.UserColumns 


	row := r.db.QueryRow(query, name, userID)
	return dbutils.UserScan(row)
}




func (r *UserRepository) UpdateAccountPassword(userID int64, password string) (bool, *errors.Error) {
	query := ` UPDATE users SET password = $1 WHERE id = $2`

	res, err := r.db.Exec(query, password, userID)
	if err != nil {
		return false, errors.Internal(err)
	}

	rowsAffected, _ := res.RowsAffected()
	return rowsAffected > 0, nil
}



