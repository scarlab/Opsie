package dbutils

import (
	"database/sql"
	"opsie/pkg/errors"
	"opsie/types"
)

const OrganizationColumns = `
    id,
    name,
    description,
    logo,
    updated_at,
    created_at
`

func OrganizationScan(row *sql.Row) (types.Organization, *errors.Error) {
	var organization types.Organization
	var logo, description  sql.NullString

	err := row.Scan(
		&organization.ID,
		&organization.Name,
		&description,
		&logo,
		&organization.UpdatedAt,
		&organization.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Organization{}, errors.NotFound("Organization not found")
		}
		return types.Organization{}, errors.Internal(err)
	}

	// Handle nullable avatar
	organization.Logo = ""
	if logo.Valid {
		organization.Logo = logo.String
	}
	organization.Description = ""
	if description.Valid {
		organization.Description = description.String
	}


	return organization, nil
}



func OrganizationScanRows(rows *sql.Rows) ([]types.Organization, *errors.Error) {
	var orgs []types.Organization

	for rows.Next() {
		var org types.Organization
		var logo, description sql.NullString

		if err := rows.Scan(
			&org.ID,
			&org.Name,
			&description,
			&logo,
			&org.UpdatedAt,
			&org.CreatedAt,
		); err != nil {
			return nil, errors.Internal(err)
		}

		if logo.Valid {
			org.Logo = logo.String
		}
		if description.Valid {
			org.Description = description.String
		}

		orgs = append(orgs, org)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Internal(err)
	}

	return orgs, nil
}
