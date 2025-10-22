package repo

import (
	"context"
	"database/sql"
	"opsie/core/dbutils"
	"opsie/pkg/errors"
	"opsie/types"
)

// internal/repository/user_organization.go
type UserOrganizationRepository struct {
    db *sql.DB
}

// NewUserOrganizationRepository creates a new instance of UserOrganizationRepository.
func NewUserOrganizationRepository(db *sql.DB) *UserOrganizationRepository {
	return &UserOrganizationRepository{db: db}
}

func (r *UserOrganizationRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, *errors.Error) {
    tx, err := r.db.BeginTx(ctx, opts)
    if err != nil {
        return nil, errors.Internal(err)
    }
    return tx, nil
}

func (r *UserOrganizationRepository) AddUserToOrg(tx *sql.Tx, userID, orgID types.ID, isDefault bool, invitedBy *types.ID) *errors.Error {
	query := `
		INSERT INTO user_organizations (user_id, organization_id, invited_by, is_default)
		VALUES ($1, $2, $3, $4)
	`

	var err error
	if tx != nil {
		_, err = tx.Exec(query, userID, orgID, invitedBy, isDefault)
	} else {
		_, err = r.db.Exec(query, userID, orgID, invitedBy, isDefault)
	}

	if err != nil {
		return errors.Internal(err)
	}

	return nil
}


func (r *UserOrganizationRepository) RemoveUserFromOrg(userID, orgID types.ID) *errors.Error {
    query := `
        DELETE FROM user_organizations
        WHERE user_id = $1 AND organization_id = $2
    `
    _, err := r.db.Exec(query, userID, orgID)
    if err != nil {
        return errors.Internal(err)
    }
    return nil
}

func (r *UserOrganizationRepository) ListOrgsByUser(userID types.ID) ([]types.UserOrganization, *errors.Error) {
	query := `
        SELECT ` + dbutils.UserOrganizationColumns + `
        FROM organizations o
        JOIN user_organizations uo ON o.id = uo.organization_id
        WHERE uo.user_id = $1
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, errors.Internal(err)
	}
	defer rows.Close()

	return dbutils.UserOrganizationScanRows(rows)
}

func (r *UserOrganizationRepository) DefaultOrg(userID types.ID) (types.UserOrganization, *errors.Error) {
    query := `
        SELECT ` + dbutils.UserOrganizationColumns + `
        FROM organizations o
        JOIN user_organizations uo ON o.id = uo.organization_id
        WHERE uo.user_id = $1 AND uo.id_default = true
        LIMIT 1
    `

    row := r.db.QueryRow(query, userID)

    return dbutils.UserOrganizationScan(row)
}



func (r *UserOrganizationRepository) SetDefaultOrg(userID, orgID types.ID) *errors.Error {
	tx, err := r.db.Begin()
	if err != nil {
		return errors.Internal(err)
	}
	defer tx.Rollback() // safe rollback if commit fails

	// Unset existing default
	_, err = tx.Exec(`UPDATE user_organizations SET is_default = false WHERE user_id = $1 AND is_default = true`, userID)
	if err != nil {
		return errors.Internal(err)
	}

	// Set new default
	_, err = tx.Exec(`UPDATE user_organizations SET is_default = true WHERE user_id = $1 AND organization_id = $2`, userID, orgID)
	if err != nil {
		return errors.Internal(err)
	}

	if err := tx.Commit(); err != nil {  // <-- wrap standard error
		return errors.Internal(err)
	}

	return nil
}

