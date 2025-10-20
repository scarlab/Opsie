package repo

import (
	"context"
	"database/sql"
	"opsie/pkg/errors"
	"opsie/types"
)

// {{.Name}}Repository - Handles DB operations for {{.Name}}.
// Talks only to the database (or other data sources).
type {{.Name}}Repository struct {
	db *sql.DB
}

// New{{.Name}}Repository - Constructor for {{.Name}}Repository
func New{{.Name}}Repository(db *sql.DB) *{{.Name}}Repository {
	return &{{.Name}}Repository{
		db: db,
	}
}

func (r *{{.Name}}Repository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, *errors.Error) {
    tx, err := r.db.BeginTx(ctx, opts)
    if err != nil {
        return nil, errors.Internal(err)
    }
    return tx, nil
}

func (r *{{.Name}}Repository) Create() (types.{{.Name}}, *errors.Error) {
    
	return types.{{.Name}}{}, errors.BadRequest("Error in repo")
}
