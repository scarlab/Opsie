package {{.PackageName}}

import (
	"net/http"
	"opsie/core/services"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/types"
	"opsie/core/socket"
)

// {{.Name}} Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *services.{{.Name}}Service
	socketHub *socket.Hub
}

// NewHandler - Constructor for {{.Name}} Handler
func NewHandler(service *services.{{.Name}}Service, socketHub *socket.Hub) *Handler {
	return &Handler{
		service: service,
		socketHub: socketHub,
	}
}



func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload types.New{{.Name}}Payload
	bolt.ParseBody(w, r, &payload)

	// Create {{.Name}}
	{{.PackageName}}, err := h.service.Create(payload)
	if err != nil {
		return err
	}

   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"	: "{{.Name}} created",
		"{{.PackageName}}"		: {{.PackageName}},
	})n nil
}