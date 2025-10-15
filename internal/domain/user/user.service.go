package user

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

// Example method:
// func (s *Service) getSomething() ([]Item, error) {
//     return s.repo.fetchSomething()
// }
