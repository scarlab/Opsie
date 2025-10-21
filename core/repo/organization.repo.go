package repo

import (
	"context"
	"database/sql"
	"opsie/core/dbutils"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"

	"github.com/lib/pq"
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
		INSERT INTO organizations (id, name, slug, description, logo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING ` + dbutils.OrganizationColumns + `;
	`

	var row *sql.Row
	ID := utils.GenerateID()
	slug := utils.Text.Slugify(payload.Name)

	if tx != nil {
		row = tx.QueryRow(query, ID, payload.Name, slug, payload.Description, payload.Logo)
	} else {
		row = r.db.QueryRow(query, ID, payload.Name, slug, payload.Description, payload.Logo)
	}
	org, err := dbutils.OrganizationScan(row)
	if err != nil {
		// Handle unique constraint violation
		if pqErr, ok := err.Original().(*pq.Error); ok && pqErr.Code == "23505" {
			return types.Organization{}, errors.Conflict("Organization already exist")
		}
		return types.Organization{}, err
	}

	return org, err
}


