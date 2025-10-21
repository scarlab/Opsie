package dbutils

import (
	"database/sql"
	"opsie/pkg/errors"
	"opsie/types"
)


const UserOrganizationColumns = `
    o.id,
    o.name,
    o.description,
    o.logo,
    o.updated_at,
    o.created_at,
    uo.is_default,
    uo.joined_at
`

func UserOrganizationScan(row *sql.Row) (types.UserOrganization, *errors.Error) {
	var org types.UserOrganization
	var logo, description sql.NullString
	var isDefault sql.NullBool
	var joinedAt sql.NullTime

	err := row.Scan(
		&org.ID,
		&org.Name,
		&description,
		&logo,
		&org.UpdatedAt,
		&org.CreatedAt,
		&isDefault,
		&joinedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.UserOrganization{}, errors.NotFound("Organization not found")
		}
		return types.UserOrganization{}, errors.Internal(err)
	}

	if logo.Valid {
		org.Logo = logo.String
	}
	if description.Valid {
		org.Description = description.String
	}
	org.IsDefault = isDefault.Valid && isDefault.Bool

	return org, nil
}

func UserOrganizationScanRows(rows *sql.Rows) ([]types.UserOrganization, *errors.Error) {
	var orgs []types.UserOrganization

	for rows.Next() {
		var org types.UserOrganization
		var logo, description sql.NullString
		var isDefault sql.NullBool
		var joinedAt sql.NullTime

		if err := rows.Scan(
			&org.ID,
			&org.Name,
			&description,
			&logo,
			&org.UpdatedAt,
			&org.CreatedAt,
			&isDefault,
			&joinedAt,
		); err != nil {		

			return nil, errors.Internal(err)
		}

		if logo.Valid {
			org.Logo = logo.String
		}
		if description.Valid {
			org.Description = description.String
		}
		org.IsDefault = isDefault.Valid && isDefault.Bool

		orgs = append(orgs, org)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Internal(err)
	}

	return orgs, nil
}
