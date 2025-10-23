package repo

import (
	"context"
	"database/sql"
	"opsie/core/dbutils"
	"opsie/core/models"
	"opsie/pkg/errors"

	"gorm.io/gorm"
)

// internal/repository/user_team.go
type UserTeamRepository struct {
    db *gorm.DB
}

// NewUserTeamRepository creates a new instance of UserTeamRepository.
func NewUserTeamRepository(db *gorm.DB) *UserTeamRepository {
	return &UserTeamRepository{db: db}
}

func (r *UserTeamRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, *errors.Error) {
    tx, err := r.db.BeginTx(ctx, opts)
    if err != nil {
        return nil, errors.Internal(err)
    }
    return tx, nil
}

func (r *UserTeamRepository) AddUserToTeam(tx *sql.Tx, userID, teamID int64, isDefault bool, invitedBy *int64) *errors.Error {
	query := `
		INSERT INTO user_teams (user_id, team_id, invited_by, is_default)
		VALUES ($1, $2, $3, $4)
	`

	var err error
	if tx != nil {
		_, err = tx.Exec(query, userID, teamID, invitedBy, isDefault)
	} else {
		_, err = r.db.Exec(query, userID, teamID, invitedBy, isDefault)
	}

	if err != nil {
		return errors.Internal(err)
	}

	return nil
}


func (r *UserTeamRepository) RemoveUserFromTeam(userID, teamID int64) *errors.Error {
    query := `
        DELETE FROM user_teams
        WHERE user_id = $1 AND team_id = $2
    `
    _, err := r.db.Exec(query, userID, teamID)
    if err != nil {
        return errors.Internal(err)
    }
    return nil
}

func (r *UserTeamRepository) ListTeamsByUser(userID int64) ([]models.UserTeam, *errors.Error) {
	query := `
        SELECT ` + dbutils.UserTeamColumns + `
        FROM teams o
        JOIN user_teams uo ON o.id = uo.team_id
        WHERE uo.user_id = $1
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.Internal(err)
	}
	defer rows.Close()

	return dbutils.UserTeamScanRows(rows)
}

func (r *UserTeamRepository) DefaultTeam(userID int64) (models.UserTeam, *errors.Error) {
    query := `
        SELECT ` + dbutils.UserTeamColumns + `
        FROM teams o
        JOIN user_teams uo ON o.id = uo.team_id
        WHERE uo.user_id = $1 AND uo.id_default = true
        LIMIT 1
    `

    row := r.db.QueryRow(query, userID)

    return dbutils.UserTeamScan(row)
}



func (r *UserTeamRepository) SetDefaultTeam(userID, teamID int64) *errors.Error {
	tx, err := r.db.Begin()
	if err != nil {
		return errors.Internal(err)
	}
	defer tx.Rollback() // safe rollback if commit fails

	// Unset existing default
	_, err = tx.Exec(`UPDATE user_teams SET is_default = false WHERE user_id = $1 AND is_default = true`, userID)
	if err != nil {
		return errors.Internal(err)
	}

	// Set new default
	_, err = tx.Exec(`UPDATE user_teams SET is_default = true WHERE user_id = $1 AND team_id = $2`, userID, teamID)
	if err != nil {
		return errors.Internal(err)
	}

	if err := tx.Commit(); err != nil {  // <-- wrap standard error
		return errors.Internal(err)
	}

	return nil
}

