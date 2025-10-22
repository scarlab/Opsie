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

// TeamRepository - Handles DB operations for Team.
// Talks only to the database (or other data sources).
type TeamRepository struct {
	db *sql.DB
}

// NewTeamRepository - Constructor for TeamRepository
func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (r *TeamRepository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, *errors.Error) {
    tx, err := r.db.BeginTx(ctx, opts)
    if err != nil {
        return nil, errors.Internal(err)
    }
    return tx, nil
}

func (r *TeamRepository) Create(tx *sql.Tx, payload types.NewTeamPayload) (types.Team, *errors.Error) {
	query := `
		INSERT INTO teams (id, name, slug, description, logo)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING ` + dbutils.TeamColumns + `;
	`

	var row *sql.Row
	ID := utils.GenerateID()
	slug := utils.Text.Slugify(payload.Name)

	if tx != nil {
		row = tx.QueryRow(query, ID, payload.Name, slug, payload.Description, payload.Logo)
	} else {
		row = r.db.QueryRow(query, ID, payload.Name, slug, payload.Description, payload.Logo)
	}
	team, err := dbutils.TeamScan(row)
	if err != nil {
		// Handle unique constraint violation
		if pqErr, ok := err.Original().(*pq.Error); ok && pqErr.Code == "23505" {
			return types.Team{}, errors.Conflict("Team already exist")
		}
		return types.Team{}, err
	}

	return team, err
}


