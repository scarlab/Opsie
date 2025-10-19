package {{.PackageName}}

import (
	"encoding/json"
	"net/http"
	"time"
	"opsie/core/services"
)

// {{.Name}} Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *{{.Name}}Service
}

// NewHandler - Constructor for {{.Name}} Handler
func NewHandler(service *{{.Name}}Service) *Handler {
	return &Handler{
		service: service,
	}
}


// func (h *Handler) Example(w http.ResponseWriter, r *http.Request) {
//     
//     json.NewEncoder(w).Encode("ok")
// }