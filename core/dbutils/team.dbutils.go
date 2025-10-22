package dbutils

import (
	"database/sql"
	"opsie/pkg/errors"
	"opsie/types"
)

const TeamColumns = `
    id,
    name,
    description,
    logo,
    updated_at,
    created_at
`


func TeamScan(row *sql.Row) (types.Team, *errors.Error) {
	var team types.Team
	var logo, description sql.NullString

	err := row.Scan(
		&team.ID,
		&team.Name,
		&description,
		&logo,
		&team.UpdatedAt,
		&team.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Team{}, errors.NotFound("Team not found")
		}
		return types.Team{}, errors.Internal(err)
	}

	if logo.Valid {
		team.Logo = logo.String
	}
	if description.Valid {
		team.Description = description.String
	}
	
	return team, nil
}

func TeamScanRows(rows *sql.Rows) ([]types.Team, *errors.Error) {
	var teams []types.Team

	for rows.Next() {
		var team types.Team
		var logo, description sql.NullString

		if err := rows.Scan(
			&team.ID,
			&team.Name,
			&description,
			&logo,
			&team.UpdatedAt,
			&team.CreatedAt,
		); err != nil {		

			return nil, errors.Internal(err)
		}

		if logo.Valid {
			team.Logo = logo.String
		}
		if description.Valid {
			team.Description = description.String
		}

		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Internal(err)
	}

	return teams, nil
}
