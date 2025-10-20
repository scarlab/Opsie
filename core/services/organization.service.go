package services

import (
	"opsie/core/repo"
	"opsie/pkg/errors"
	"opsie/types"
)

// OrganizationService - Contains all business logic for Organization api.
// Talks to the Repository, but never to HTTP directly.
type OrganizationService struct {
	repo *repo.OrganizationRepository
}

// NewOrganizationService - Constructor for OrganizationService
func NewOrganizationService(repo *repo.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		repo: repo,
	}
}


func (s *OrganizationService) Create(payload types.NewOrganizationPayload) (types.Organization, *errors.Error) {
	if payload.Name == "" {
		return types.Organization{}, errors.BadRequest("Organization name ir required")
	}

	organization, err := s.repo.Create(nil, payload)
	if err != nil {
		return types.Organization{}, err
	}

    return organization, nil
}
