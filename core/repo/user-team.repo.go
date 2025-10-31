package repo

import (
	"opsie/core/models"
	"opsie/pkg/errors"

	"gorm.io/gorm"
)

type UserTeamRepository struct {
	db *gorm.DB
}

func NewUserTeamRepository(db *gorm.DB) *UserTeamRepository {
	return &UserTeamRepository{db: db}
}

// AddUserToTeam adds a user to a team
func (r *UserTeamRepository) AddUserToTeam(payload models.AddUserToTeamPayload) *errors.Error {
	ut := models.UserTeam{
		UserID:    payload.UserID,
		TeamID:    payload.TeamID,
		IsDefault: payload.IsDefault,
		IsAdmin:   payload.IsAdmin,
		InvitedBy: payload.InvitedBy,
	}

	if err := r.db.Create(&ut).Error; err != nil {
		return errors.Internal(err)
	}
	return nil
}



// RemoveUserFromTeam removes a user from a team
func (r *UserTeamRepository) RemoveUserFromTeam(userID, teamID int64) *errors.Error {
	if err := r.db.Where("user_id = ? AND team_id = ?", userID, teamID).Delete(&models.UserTeam{}).Error; err != nil {
		return errors.Internal(err)
	}
	return nil
}


// RemoveAllUserFromTeam: removes every user from the team
func (r *UserTeamRepository) RemoveAllUserFromTeam(teamID int64) *errors.Error {
    if err := r.db.Where("team_id = ?", teamID).Delete(&models.UserTeam{}).Error; err != nil {
        return errors.Internal(err)
    }
    return nil
}


// ListTeamsByUser returns all teams a user belongs to
func (r *UserTeamRepository) ListTeamsByUser(userID int64) ([]models.TeamWithMeta, *errors.Error) {
	var teams []models.TeamWithMeta
	
	err := r.db.
		Table("teams").
		Select("teams.*, ut.is_default, ut.is_admin").
		Joins("JOIN user_teams ut ON ut.team_id = teams.id").
		Where("ut.user_id = ?", userID).
		Scan(&teams).Error

	if err != nil {
		return nil, errors.Internal(err)
	}

	return teams, nil
}

// DefaultTeam returns the default team for a user along with metadata
func (r *UserTeamRepository) DefaultTeam(userID int64) (models.TeamWithMeta, *errors.Error) {
	var teamMeta models.TeamWithMeta

	err := r.db.
		Table("teams").
		Select("teams.*, ut.is_default, ut.is_admin").
		Joins("JOIN user_teams ut ON ut.team_id = teams.id").
		Where("ut.user_id = ? AND ut.is_default = true", userID).
		Scan(&teamMeta).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.TeamWithMeta{}, errors.NotFound("default team not found")
		}
		return models.TeamWithMeta{}, errors.Internal(err)
	}

	return teamMeta, nil
}

// SetDefaultTeam sets a user's default team and returns the updated team (transaction-safe)
func (r *UserTeamRepository) SetDefaultTeam(userID, teamID int64) (*models.Team, *errors.Error) {
	var team models.Team

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Unset existing default
		if err := tx.Model(&models.UserTeam{}).
			Where("user_id = ? AND is_default = true", userID).
			Update("is_default", false).Error; err != nil {
			return err
		}

		// Set new default
		if err := tx.Model(&models.UserTeam{}).
			Where("user_id = ? AND team_id = ?", userID, teamID).
			Update("is_default", true).Error; err != nil {
			return err
		}

		// Fetch and return the team details (optional join for richer data)
		if err := tx.Model(&models.Team{}).
			Where("id = ?", teamID).
			First(&team).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errors.Internal(err)
	}

	return &team, nil
}

