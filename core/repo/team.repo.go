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
            ID: utils.GenerateID(), // snowflake ID
        },
        Name:        payload.Name,
        Slug:        utils.Text.Slugify(payload.Name),
        Description: payload.Description,
        Logo:        payload.Logo,
    }

    if err := r.db.Create(&team).Error; err != nil {
        // Handle unique constraint violation (slug or name)
        if errors.IsPgConflict(err) { 
            return models.Team{}, errors.Conflict("Team already exists")
        }
        return models.Team{}, errors.Internal(err)
    }

    return team, nil
}


