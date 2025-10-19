package {{.PackageName}}

import (
	"encoding/json"
	"net/http"
	"opsie/core/socket"
	"opsie/core/services"
	"time"
)

// {{.Name}} Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *{{.Name}}Service
	socketHub *socket.Hub
}

// NewHandler - Constructor for {{.Name}} Handler
func NewHandler(service *{{.Name}}Service, socketHub *socket.Hub) *Handler {
	return &Handler{
		service: service,
		socketHub: socketHub,
	}
}


// func (h *Handler) Example(w http.ResponseWriter, r *http.Request) {
//     
//     json.NewEncoder(w).Encode("ok")
// }
