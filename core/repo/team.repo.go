package repo

import (
	"opsie/core/models"
	"opsie/pkg/errors"
	"opsie/pkg/utils"

	"gorm.io/gorm"
)

// TeamRepository - Handles DB operations for Team.
// Talks only to the database (or other data sources).
type TeamRepository struct {
	db *gorm.DB
}

// NewTeamRepository - Constructor for TeamRepository
func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}


func (r *TeamRepository) Create(payload models.NewTeamPayload) (models.Team, *errors.Error) {
    team := models.Team{
        BaseModel: models.BaseModel{
            ID: utils.GenerateID(),
        },
        Name:        payload.Name,
        Slug:        utils.Text.Slugify(payload.Name),
        Description: payload.Description,
    }

    if err := r.db.Create(&team).Error; err != nil {
        // Handle unique constraint violation (slug or name)
        if errors.IsPgConflict(err) {
            return models.Team{}, errors.Conflict("Team with this name already exists")
        }
        return models.Team{}, errors.Internal(err)
    }

    return team, nil
}


func (r *TeamRepository) Count() (int64, *errors.Error) {
    var count int64
    if err := r.db.Model(&models.Team{}).Count(&count).Error; err != nil {
        return 0, errors.Internal(err)
    }
    return count, nil
}

func (r *TeamRepository) GetAll() ([]models.Team, *errors.Error) {
    var teams []models.Team

    if err := r.db.Find(&teams).Error; err != nil {
        return nil, errors.Internal(err)
    }

    return teams, nil
}


func (r *TeamRepository) GetById(id int64) (models.Team, *errors.Error) {
	var team models.Team

	if err := r.db.Where("id = ?", id).First(&team).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Team{}, errors.NotFound("Team not found")
		}
		return models.Team{}, errors.Internal(err)
	}

	return team, nil
}



func (r *TeamRepository) GetAllByUserId(userId int64) ([]models.Team, *errors.Error) {
    var team []models.Team

    if err := r.db.Find(&team).Where("user_id = ?", userId).Error; err != nil {
        return [] models.Team{}, errors.Internal(err)
    }

    return team, nil
}


// Update updates an existing team by ID with the provided payload
func (r *TeamRepository) Update(id int64, payload models.NewTeamPayload) (models.Team, *errors.Error) {
	var team models.Team

	// Find existing team
	if err := r.db.First(&team, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
            return models.Team{}, errors.NotFound("Team not found")
        }
		return models.Team{}, errors.Internal(err)
	}

    // Update fields
	team.Name = payload.Name
	team.Description = payload.Description

	// Regenerate slug if name changed
	if payload.Name != "" {
		team.Slug = utils.Text.Slugify(payload.Name)
	}

	// Save updates
	if err := r.db.Save(&team).Error; err != nil {
		if errors.IsPgConflict(err) {
			return models.Team{}, errors.Conflict("Team with this name already exists")
		}
		return models.Team{}, errors.Internal(err)
	}

	return team, nil
}



// Delete removes a team by ID
func (r *TeamRepository) Delete(id int64) *errors.Error {
	result := r.db.Delete(&models.Team{}, id)
	if result.Error != nil {
		return  errors.Internal(result.Error)
	}
	if result.RowsAffected == 0 {
		return  errors.NotFound("User not found")
	}
	return nil
}




