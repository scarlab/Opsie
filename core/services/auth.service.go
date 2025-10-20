package services

import (
	"opsie/config"
	"opsie/core/repo"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"
	"time"
)

// AuthService - Contains all business logic for this api.
// Talks to the Repository, but never to HTTP directly.
type AuthService struct {
	repo *repo.AuthRepository
	userRepo *repo.UserRepository
}

// NewService - Constructor for Service
func NewAuthService(repo *repo.AuthRepository, userRepo *repo.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
		userRepo: userRepo,
	}
}


func (s *AuthService) AuthenticateUser(payload types.LoginPayload) (types.AuthUser, *errors.Error) {
	// Basic validation
	if payload.Email == "" || payload.Password == "" {
		return types.AuthUser{}, errors.BadRequest("email and password required")
	}
	
	
	// Get Request User By Email
	reqUser, err := s.userRepo.GetByEmail(payload.Email)
	if err != nil {
		return types.AuthUser{}, err
	}

	// Compare Password
	isMatched := utils.Hash.Compare(reqUser.Password, payload.Password)
	
	if !isMatched {
	  return  types.AuthUser{}, errors.Unauthorized("invalid email or password")
	}

	// Generate Auth user
	authUser := types.AuthUser{
		ID: reqUser.ID,
		DisplayName: reqUser.DisplayName,
		Email: reqUser.Email,
		Avatar: reqUser.Avatar,
		SystemRole: reqUser.SystemRole,
		Preference: reqUser.Preference,
	}
	return authUser, nil
}


func (s *AuthService) CreateSession(userID types.ID) (types.Session, *errors.Error) {
	key, err := utils.GenerateSessionKey()
	if err != nil {
		errors.Internal(err)
	}

	expiry := time.Now().Add(time.Duration(config.AppConfig.SessionDays) * 24 * time.Hour)

	session, err1 := s.repo.CreateSession(userID, key, expiry)
	if err1 != nil {
		return types.Session{}, err1
	}

	return session, nil
}



func (s *AuthService) HandleLogout(key types.SessionKey) *errors.Error {
	// 1. Expire the session in DB
	queryErr := s.repo.ExpireSession(string(key))
	if queryErr != nil {
		return errors.Internal(queryErr.Err)
	}
	return nil
}
