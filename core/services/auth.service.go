package services

import (
	"opsie/config"
	"opsie/core/models"
	"opsie/core/repo"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
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


func (s *AuthService) AuthenticateUser(payload models.LoginPayload) (models.AuthUser, *errors.Error) {
	// Basic validation
	if payload.Email == "" || payload.Password == "" {
		return models.AuthUser{}, errors.BadRequest("email and password required")
	}
	
	
	// Get Request User By Email
	reqUser, err := s.userRepo.GetByEmail(payload.Email)
	if err != nil {
		return models.AuthUser{}, err
	}

	// Compare Password
	isMatched := utils.Hash.Compare(reqUser.Password, payload.Password)
	
	if !isMatched {
	  return  models.AuthUser{}, errors.Unauthorized("invalid email or password")
	}

	// Generate Auth user
	authUser := models.AuthUser{
		ID: reqUser.ID,
		DisplayName: reqUser.DisplayName,
		Email: reqUser.Email,
		Avatar: reqUser.Avatar,
		SystemRole: reqUser.SystemRole,
		Preference: reqUser.Preference,
	}
	return authUser, nil
}


func (s *AuthService) CreateSession(userID int64) (models.Session, *errors.Error) {
	key, err := utils.GenerateSessionKey()
	if err != nil {
		errors.Internal(err)
	}

	expiry := time.Now().Add(time.Duration(config.App.SessionDays) * 24 * time.Hour)

	session, err1 := s.repo.CreateSession(userID, key, expiry)
	if err1 != nil {
		return models.Session{}, err1
	}

	return session, nil
}



func (s *AuthService) HandleLogout(key string) *errors.Error {
	// 1. Expire the session in DB
	queryErr := s.repo.ExpireSession(string(key))
	if queryErr != nil {
		return errors.Internal(queryErr.Err)
	}
	return nil
}
