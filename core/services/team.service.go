package services

import (
	"opsie/core/repo"
	"opsie/pkg/errors"
	"opsie/types"
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


func (s *TeamService) Create(payload types.NewTeamPayload) (types.Team, *errors.Error) {
	if payload.Name == "" {
		return types.Team{}, errors.BadRequest("Team name is required")
	}

	team, err := s.repo.Create(nil, payload)
	if err != nil {
		return types.Team{}, err
	}

	// Add Owner 

    return team, nil
}



func (s *TeamService) GetUserTeams(userId types.ID) ([]types.UserTeam, *errors.Error) {
	if userId <= 0 {
		return []types.UserTeam{}, errors.BadRequest("User id is required")
	}

	teams, err := s.userTeamRepo.ListTeamsByUser(userId)
	if err != nil {
		return []types.UserTeam{}, err
	}

    return teams, nil
}



func (s *TeamService) GetUserDefaultTeam(userId types.ID) (types.UserTeam, *errors.Error) {
	if userId <= 0 {
		return types.UserTeam{}, errors.BadRequest("User id is required")
	}

	team, err := s.userTeamRepo.DefaultTeam(userId)
	if err != nil {
		return types.UserTeam{}, err
	}

    return team, nil
}
