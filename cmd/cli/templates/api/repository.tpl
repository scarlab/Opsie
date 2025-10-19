package repo

import "database/sql"

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


// func (r *{{.Name}}Repository) Example() (Item, error) {
//     rows, err := r.db.Query("SELECT * FROM users")
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()
//     // Map rows to struct...
// }
