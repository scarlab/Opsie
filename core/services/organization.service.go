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
	userOrgRepo *repo.UserOrganizationRepository
}

// NewOrganizationService - Constructor for OrganizationService
func NewOrganizationService(repo *repo.OrganizationRepository, userOrgRepo *repo.UserOrganizationRepository) *OrganizationService {
	return &OrganizationService{
		repo: repo,
		userOrgRepo: userOrgRepo,
	}
}


func (s *OrganizationService) Create(payload types.NewOrganizationPayload) (types.Organization, *errors.Error) {
	if payload.Name == "" {
		return types.Organization{}, errors.BadRequest("Organization name is required")
	}

	organization, err := s.repo.Create(nil, payload)
	if err != nil {
		return types.Organization{}, err
	}

    return organization, nil
}



func (s *OrganizationService) GetUserOrganizations(userId types.ID) ([]types.UserOrganization, *errors.Error) {
	if userId <= 0 {
		return []types.UserOrganization{}, errors.BadRequest("Organization id is required")
	}

	organizations, err := s.userOrgRepo.ListOrgsByUser(userId)
	if err != nil {
		return []types.UserOrganization{}, err
	}

    return organizations, nil
}
