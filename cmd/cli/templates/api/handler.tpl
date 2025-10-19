package {{.PackageName}}

import (
	"encoding/json"
	"net/http"
	"opsie/core/services"
	"opsie/pkg/errors"
)


// {{.Name}} Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *services.{{.Name}}Service
}

// NewHandler - Constructor for {{.Name}} Handler
func NewHandler(service *services.{{.Name}}Service) *Handler {
	return &Handler{
		service: service,
	}
}



func (h *Handler) Example(w http.ResponseWriter, r *http.Request) *errors.Error{
    
    json.NewEncoder(w).Encode("ok")
	return nil
}