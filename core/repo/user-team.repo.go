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
func (r *UserTeamRepository) AddUserToTeam(userID, teamID int64, isDefault bool, invitedBy *int64) *errors.Error {
	ut := models.UserTeam{
		UserID:    userID,
		TeamID:    teamID,
		InvitedBy: invitedBy,
		IsDefault: isDefault,
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

// ListTeamsByUser returns all teams a user belongs to
func (r *UserTeamRepository) ListTeamsByUser(userID int64) ([]models.UserTeam, *errors.Error) {
	var teams []models.UserTeam
	if err := r.db.Where("user_id = ?", userID).Find(&teams).Error; err != nil {
		return nil, errors.Internal(err)
	}
	return teams, nil
}

// DefaultTeam returns the default team for a user
func (r *UserTeamRepository) DefaultTeam(userID int64) (models.UserTeam, *errors.Error) {
	var team models.UserTeam
	if err := r.db.Where("user_id = ? AND is_default = true", userID).First(&team).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.UserTeam{}, errors.NotFound("default team not found")
		}
		return models.UserTeam{}, errors.Internal(err)
	}
	return team, nil
}

// SetDefaultTeam sets a user's default team (transaction-safe)
func (r *UserTeamRepository) SetDefaultTeam(userID, teamID int64) *errors.Error {
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

		return nil
	})

	if err != nil {
		return errors.Internal(err)
	}
	return nil
}
