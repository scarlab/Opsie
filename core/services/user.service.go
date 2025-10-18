package services

import (
	repo "opsie/core/repositories"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"opsie/types"
)

// UserService - Contains all business logic for this domain.
// Talks to the Repository, but never to HTTP directly.
type UserService struct {
	repo *repo.UserRepository
}

// NewService - Constructor for Service
func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{
		repo: repo,
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

