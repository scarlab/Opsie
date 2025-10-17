package user

import (
	"opsie/pkg/errors"
	"opsie/pkg/utils"
)

// Service - Contains all business logic for this domain.
// Talks to the Repository, but never to HTTP directly.
type Service struct {
	repo *Repository
}

// NewService - Constructor for Service
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateOwnerAccount handles business logic for creating the first owner
func (s *Service) CreateOwnerAccount(payload TNewOwnerPayload) (TUser, *errors.Error) {
	// Basic validation
	if payload.Email == "" || payload.Password == "" {
		return TUser{}, errors.BadRequest("email and password required")
	}

	hashedPassword, _ := utils.Hash.Generate( payload.Password)
	payload.Password = hashedPassword

	createdUser, err := s.repo.CreateOwnerAccount(payload)
	if err != nil {
		return TUser{}, err
	}

	return createdUser, nil
}
