package services

import 	"opsie/core/repo"


// {{.Name}}Service - Contains all business logic for {{.Name}} api.
// Talks to the Repository, but never to HTTP directly.
type {{.Name}}Service struct {
	repo *repo.{{.Name}}Repository
}

// New{{.Name}}Service - Constructor for {{.Name}}Service
func New{{.Name}}Service(repo *repo.{{.Name}}Repository) *{{.Name}}Service {
	return &{{.Name}}Service{
		repo: repo,
	}
}


// func (s *{{.Name}}Service) Example() (Item, error) {
//     return s.repo.Example()
// }
