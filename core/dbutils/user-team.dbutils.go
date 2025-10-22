package dbutils

import (
	"database/sql"
	"opsie/pkg/errors"
	"opsie/types"
)


const UserTeamColumns = `
    o.id,
    o.name,
    o.description,
    o.logo,
    o.updated_at,
    o.created_at,
    uo.is_default,
    uo.joined_at
`

func UserTeamScan(row *sql.Row) (types.UserTeam, *errors.Error) {
	var team types.UserTeam
	var logo, description sql.NullString
	var isDefault sql.NullBool
	var joinedAt sql.NullTime

	err := row.Scan(
		&team.ID,
		&team.Name,
		&description,
		&logo,
		&team.UpdatedAt,
		&team.CreatedAt,
		&isDefault,
		&joinedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.UserTeam{}, errors.NotFound("Team not found")
		}
		return types.UserTeam{}, errors.Internal(err)
	}

	if logo.Valid {
		team.Logo = logo.String
	}
	if description.Valid {
		team.Description = description.String
	}
	team.IsDefault = isDefault.Valid && isDefault.Bool

	return team, nil
}

func UserTeamScanRows(rows *sql.Rows) ([]types.UserTeam, *errors.Error) {
	var teams []types.UserTeam

	for rows.Next() {
		var team types.UserTeam
		var logo, description sql.NullString
		var isDefault sql.NullBool
		var joinedAt sql.NullTime

		if err := rows.Scan(
			&team.ID,
			&team.Name,
			&description,
			&logo,
			&team.UpdatedAt,
			&team.CreatedAt,
			&isDefault,
			&joinedAt,
		); err != nil {		

			return nil, errors.Internal(err)
		}

		if logo.Valid {
			team.Logo = logo.String
		}
		if description.Valid {
			team.Description = description.String
		}
		team.IsDefault = isDefault.Valid && isDefault.Bool

		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Internal(err)
	}

	return teams, nil
}
