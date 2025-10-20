package repo

import (
	"context"
	"database/sql"
	"opsie/core/dbutils"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"
)

// OrganizationRepository - Handles DB operations for Organization.
// Talks only to the database (or other data sources).
type OrganizationRepository struct {
	db *sql.DB
}

// NewOrganizationRepository - Constructor for OrganizationRepository
func NewOrganizationRepository(db *sql.DB) *OrganizationRepository {
	return &OrganizationRepository{
		db: db,
	}
}

func (r *OrganizationRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, *errors.Error) {
    tx, err := r.db.BeginTx(ctx, opts)
    if err != nil {
        return nil, errors.Internal(err)
    }
    return tx, nil
}

func (r *OrganizationRepository) Create(tx *sql.Tx, payload types.NewOrganizationPayload) (types.Organization, *errors.Error) {
	query := `
		INSERT INTO organizations (id, name, description, logo)
		VALUES ($1, $2, $3, $4)
		RETURNING ` + dbutils.OrganizationColumns + `;
	`

	var row *sql.Row
	ID := utils.GenerateID()

	if tx != nil {
		row = tx.QueryRow(query, ID, payload.Name, payload.Description, payload.Logo)
	} else {
		row = r.db.QueryRow(query, ID, payload.Name, payload.Description, payload.Logo)
	}

	return dbutils.OrganizationScan(row)
}


