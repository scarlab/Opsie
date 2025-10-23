package services

import (
	"opsie/core/models"
	"opsie/core/repo"
	"opsie/pkg/errors"
)

// TeamService - Contains all business logic for Team api.
// Talks to the Repository, but never to HTTP directly.
type TeamService struct {
	repo *repo.TeamRepository
	userTeamRepo *repo.UserTeamRepository
}

// NewTeamService - Constructor for TeamService
func NewTeamService(repo *repo.TeamRepository, userTeamRepo *repo.UserTeamRepository) *TeamService {
	return &TeamService{
		repo: repo,
		userTeamRepo: userTeamRepo,
	}
}


func (s *TeamService) Create(payload models.NewTeamPayload) (models.Team, *errors.Error) {
	if payload.Name == "" {
		return models.Team{}, errors.BadRequest("Team name is required")
	}

	team, err := s.repo.Create(nil, payload)
	if err != nil {
		return models.Team{}, err
	}

	// Add Owner 

    return team, nil
}



func (s *TeamService) GetUserTeams(userId int64) ([]models.UserTeam, *errors.Error) {
	if userId <= 0 {
		return []models.UserTeam{}, errors.BadRequest("User id is required")
	}

	teams, err := s.userTeamRepo.ListTeamsByUser(userId)
	if err != nil {
		return []models.UserTeam{}, err
	}

    return teams, nil
}



func (s *TeamService) GetUserDefaultTeam(userId int64) (models.UserTeam, *errors.Error) {
	if userId <= 0 {
		return models.UserTeam{}, errors.BadRequest("User id is required")
	}

	team, err := s.userTeamRepo.DefaultTeam(userId)
	if err != nil {
		return models.UserTeam{}, err
	}

    return team, nil
}
