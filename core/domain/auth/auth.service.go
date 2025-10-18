package auth

import (
	"opsie/config"
	"opsie/core/domain/user"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"time"
)

// Service - Contains all business logic for this domain.
// Talks to the Repository, but never to HTTP directly.
type Service struct {
	repo *Repository
	userRepo *user.Repository
}

// NewService - Constructor for Service
func NewService(repo *Repository, userRepo *user.Repository) *Service {
	return &Service{
		repo: repo,
		userRepo: userRepo,
	}
}


func (s *Service) AuthenticateUser(payload TLoginPayload) (TAuthUser, *errors.Error) {
	// Basic validation
	if payload.Email == "" || payload.Password == "" {
		return TAuthUser{}, errors.BadRequest("email and password required")
	}
	
	
	// Get Request User By Email
	reqUser, err := s.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return TAuthUser{}, err
	}

	// Compare Password
	isMatched := utils.Hash.Compare(reqUser.Password, payload.Password)
	
	if !isMatched {
		errors.Unauthorized("invalid email or password")
	}

	// Generate Auth user
	authUser := TAuthUser{
		ID: reqUser.ID,
		DisplayName: reqUser.DisplayName,
		Email: reqUser.Email,
		Avatar: reqUser.Avatar,
		SystemRole: reqUser.SystemRole,
		Preference: reqUser.Preference,
	}
	return authUser, nil
}


func (s *Service) CreateSession(userID int64) (TSession, *errors.Error) {
	// 
	key, err := utils.GenerateSessionKey()
	if err != nil {
		errors.Internal(err)
	}

	expiry := time.Now().Add(time.Duration(config.AppConfig.SessionDays) * 24 * time.Hour)

	session, err1 := s.repo.CreateSession(userID, key, expiry)
	if err1 != nil {
		return TSession{}, err1
	}

	return session, nil
}
