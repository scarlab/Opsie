package {{.PackageName}}

import "database/sql"

// Repository - Handles DB operations for {{.PackageName}}.
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

// Example method:
// func (r *Repository) fetchSomething() ([]Item, error) {
//     rows, err := r.db.Query("SELECT * FROM something")
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()
//     // Map rows to struct...
// }
