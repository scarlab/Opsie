package services

import (
	"opsie/core/repo"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"
)

// UserService - Contains all business logic for this api.
// Talks to the Repository, but never to HTTP directly.
type UserService struct {
	repo *repo.UserRepository
	authRepo *repo.AuthRepository
}

// NewService - Constructor for Service
func NewUserService(repo *repo.UserRepository, authRepo *repo.AuthRepository) *UserService {
	return &UserService{
		repo: repo,
		authRepo: authRepo,
	}
}

// CreateOwnerAccount handles business logic for creating the first owner
func (s *UserService) CreateOwnerAccount(payload types.NewOwnerPayload) (types.User, *errors.Error) {
	// Basic validation
	if payload.Email == "" || payload.Password == "" {
		return types.User{}, errors.BadRequest("email and password required")
	}

	hashedPassword, _ := utils.Hash.Generate( payload.Password)
	payload.Password = hashedPassword

	createdUser, err := s.repo.CreateOwnerAccount(payload)
	if err != nil {
		return types.User{}, err
	}

	return createdUser, nil
}


// GetOwnerCount
func (s *UserService) GetOwnerCount() (int, *errors.Error) {
	count, err := s.repo.GetOwnerCount()
	if err != nil {
		return 0, err
	}

	return count, nil
}



func (s *UserService) UpdateAccountName(userID types.ID, payload types.UpdateAccountNamePayload) (types.AuthUser, *errors.Error) {
	if userID == 0 || payload.DisplayName == "" {
		return types.AuthUser{}, errors.BadRequest("Invalid user or name")
	}
	
	user, err := s.repo.UpdateAccountName(userID, payload.DisplayName)
	if err != nil {
		return types.AuthUser{}, err
	}

	
	authUser := types.AuthUser{
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



func (s *UserService) UpdateAccountPassword(userID types.ID, sessionKey types.SessionKey, payload types.UpdateAccountPasswordPayload) (types.Session, *errors.Error) {
	if userID == 0 || payload.Password == "" ||payload.NewPassword == "" {
		return types.Session{}, errors.BadRequest("Invalid user or name")
	}
	
	// Get the user from db. need password verification
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return types.Session{}, err
	}

	// Verify password
	isMatched := utils.Hash.Compare(user.Password, payload.Password)
	if !isMatched {
		return types.Session{}, errors.BadRequest("Password doesn't match")
	}

	// generate Hashed Password
	hashedPassword, _ := utils.Hash.Generate(payload.NewPassword)

	// Update the password
	_, uapErr := s.repo.UpdateAccountPassword(userID, hashedPassword)
	if uapErr != nil {
		return types.Session{}, uapErr
	}

	// Regenerate Session key
	session, rskErr := s.authRepo.RegenerateSessionKey(sessionKey)
	if rskErr != nil {
		return types.Session{}, err
	}

	return session, nil
}

