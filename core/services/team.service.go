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





func (s *TeamService) GetUserTeams(userId int64) ([]models.TeamWithMeta, *errors.Error) {
	if userId <= 0 {
		return []models.TeamWithMeta{}, errors.BadRequest("User id is required")
	}

	teams, err := s.userTeamRepo.ListTeamsByUser(userId)
	if err != nil {
		return []models.TeamWithMeta{}, err
	}

    return teams, nil
}



func (s *TeamService) GetUserDefaultTeam(userId int64) (models.TeamWithMeta, *errors.Error) {
	if userId <= 0 {
		return models.TeamWithMeta{}, errors.BadRequest("User id is required")
	}

	team, err := s.userTeamRepo.DefaultTeam(userId)
	if err != nil {
		return models.TeamWithMeta{}, err
	}

    return team, nil
}






func (s *TeamService) Create(payload models.NewTeamPayload) (models.Team, *errors.Error) {
	if payload.Name == "" {
		return models.Team{}, errors.BadRequest("Team name is required")
	}

	team, err := s.repo.Create(payload)
	if err != nil {
		return models.Team{}, err
	}

    return team, nil
}



func (s *TeamService) GetAll() ([]models.Team, *errors.Error) {
	
	teams, err := s.repo.GetAll()
	if err != nil {
		return []models.Team{}, err
	}

    return teams, nil
}


func (s *TeamService) GetById(id int64) (models.Team, *errors.Error) {
	
	team, err := s.repo.GetById(id)
	if err != nil {
		return models.Team{}, err
	}

    return team, nil
}

func (s *TeamService) GetAllByUserId(userId int64) ([]models.TeamWithMeta, *errors.Error) {
	
	team, err := s.userTeamRepo.ListTeamsByUser(userId)
	if err != nil {
		return []models.TeamWithMeta{}, err
	}

    return team, nil
}


func (s *TeamService) Update(id int64, payload models.UpdateTeamPayload) (models.Team, *errors.Error) {
	
	team, err := s.repo.Update(id, payload)
	if err != nil {
		return models.Team{}, err
	}

    return team, nil
}


func (s *TeamService) Delete(teamId int64) (bool, *errors.Error) {
	// Get total team count
	teamCount, err := s.repo.Count()
	if err != nil {return false, err}
	if teamCount <= 1 {
		return false, errors.Conflict("Cannot delete the only existing team")
	}

	// Get project count


	// Get resource count


	// Check for project & resource
	// if projectCount > 0 || resourceCount > 0 {
    //     return false, errors.Conflict("Team cannot be deleted — remove projects/resources first")
    // }


	// Remove All user from the team
	removeErr := s.userTeamRepo.RemoveAllUserFromTeam(teamId)
	if removeErr != nil {
		return false, removeErr
	}


	// Delete the team
	team, dltErr := s.repo.Delete(teamId)
	if dltErr != nil {
		return false, dltErr
	}

    return team, nil
}


