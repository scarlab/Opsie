package services

import (
	"opsie/core/models"
	"opsie/core/repo"
	"opsie/def"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
)

// UserService - Contains all business logic for this api.
// Talks to the Repository, but never to HTTP directly.
type UserService struct {
	repo *repo.UserRepository
	authRepo *repo.AuthRepository
	teamRepo *repo.TeamRepository
	userTeamRepo *repo.UserTeamRepository
}

// NewService - Constructor for Service
func NewUserService(repo *repo.UserRepository, authRepo *repo.AuthRepository,teamRepo *repo.TeamRepository, userTeamRepo *repo.UserTeamRepository) *UserService {
	return &UserService{
		repo: repo,
		authRepo: authRepo,
		teamRepo: teamRepo,
		userTeamRepo: userTeamRepo,
	}
}

/// _____________________________________________________________________________________________________________
/// Onboarding --------------------------------------------------------------------------------------------------
/// ---

// CreateOwnerAccount handles business logic for creating the first owner
func (s *UserService) CreateOwnerAccount(payload models.NewUserPayload) (models.User, *errors.Error) {
	if payload.Email == "" || payload.Password == "" {
		return models.User{}, errors.BadRequest("Email and password required")
	}

	hashedPassword, _ 	:= utils.Hash.Generate(payload.Password)
	
	payload.Password 	= hashedPassword
	payload.SystemRole 	= def.SystemRoleOwner.ToString()
	payload.ResetPass	= false


	// Create user
	user, err := s.repo.Create(payload)
	if err != nil {
		return models.User{}, err
	}

	// Create default team
	teamPayload := models.NewTeamPayload{
		Name:        utils.GenerateTeamName(),
		Description: "This is your default team.",
	}
	team, teamErr := s.teamRepo.Create( teamPayload)
	if teamErr != nil {
		return models.User{}, teamErr
	}

	// Link user <-> team
	if addErr := s.userTeamRepo.AddUserToTeam(user.ID, team.ID, true, nil, true); addErr != nil {
		return models.User{}, addErr
	}

	return user, nil
}


// GetOwnerCount
func (s *UserService) GetOwnerCount() (int, *errors.Error) {
	count, err := s.repo.GetOwnerCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}



/// _____________________________________________________________________________________________________________
/// User Account ------------------------------------------------------------------------------------------------
/// ---

func (s *UserService) UpdateAccountName(userID int64, payload models.UpdateAccountNamePayload) (models.AuthUser, *errors.Error) {
	if userID == 0 || payload.DisplayName == "" {
		return models.AuthUser{}, errors.BadRequest("Invalid user or name")
	}
	
	user, err := s.repo.UpdateAccountName(userID, payload.DisplayName)
	if err != nil {
		return models.AuthUser{}, err
	}

	
	authUser := models.AuthUser{
		ID: user.ID,
		DisplayName: user.DisplayName,
		Email: user.Email,
		Avatar: user.Avatar,
		SystemRole: user.SystemRole,
		IsActive: user.IsActive,
		Preference: user.Preference,
	}

	return authUser, nil
}



func (s *UserService) UpdateAccountPassword(userID int64, sessionKey string, payload models.UpdateAccountPasswordPayload) (models.Session, *errors.Error) {
	if userID == 0 || payload.Password == "" ||payload.NewPassword == "" {
		return models.Session{}, errors.BadRequest("Invalid user or name")
	}
	
	// Get the user from db. need password verification
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return models.Session{}, err
	}

	// Verify password
	isMatched := utils.Hash.Compare(user.Password, payload.Password)
	if !isMatched {
		return models.Session{}, errors.BadRequest("Password doesn't match")
	}

	// generate Hashed Password
	hashedPassword, _ := utils.Hash.Generate(payload.NewPassword)

	// Update the password
	_, uapErr := s.repo.UpdateAccountPassword(userID, hashedPassword)
	if uapErr != nil {
		return models.Session{}, uapErr
	}

	// Regenerate Session key
	session, rskErr := s.authRepo.RegenerateSessionKey(sessionKey)
	if rskErr != nil {
		return models.Session{}, err
	}

	return session, nil
}





/// _____________________________________________________________________________________________________________
/// Admin Only --------------------------------------------------------------------------------------------------
/// ---

func (s *UserService) CreateUser(payload models.NewUserPayload) (models.User, *errors.Error) {
	if payload.Email == "" || payload.Password == "" {
		return models.User{}, errors.BadRequest("Email and password required")
	}

	hashedPassword, _ := utils.Hash.Generate(payload.Password)
	payload.Password = hashedPassword

	// Create user
	user, err := s.repo.Create(payload)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}



func (s *UserService) GetAllUser() ([]models.User, *errors.Error) {
	
	// Create user
	users, err := s.repo.GetAll()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}



func (s *UserService) GetUserById(id int64) (models.User, *errors.Error) {
	
	// get the user of ID
	user, err := s.repo.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}



func (s *UserService) UpdateUser(id int64, payload models.UpdateUserPayload) (models.User, *errors.Error) {
	

	return models.User{}, nil
}



func (s *UserService) DeleteUser(id int64,)  *errors.Error {
	// Delete the user
	return s.repo.Delete(id)
}



func (s *UserService) AddUserToTeam(userId, teamId int64) *errors.Error {
	

	return nil
}



func (s *UserService) RemoveUserFromTeam(userId, teamId int64) *errors.Error {
	

	return  nil
}