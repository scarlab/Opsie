package services

import 	(
	"opsie/core/repo"
	"opsie/pkg/errors"
	"opsie/types"
	"time"
)



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


func (s *{{.Name}}Service) Create(payload types.New{{.Name}}Payload) (types.{{.Name}}, *errors.Error) {
	if payload.Name == "" {
		return types.{{.Name}}{}, errors.BadRequest("{{.Name}} name ir required")
	}

    return types.{{.Name}}{Name: payload.Name, CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil
}
