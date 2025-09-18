package {{.PackageName}}

import (
	"encoding/json"
	"net/http"
	"time"
)

// Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *Service
}

// NewHandler - Constructor for Handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Example method:
// func (h *Handler) getSomething(w http.ResponseWriter, r *http.Request) {
//     result, err := h.service.getSomething()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     json.NewEncoder(w).Encode(result)
// }



// Health checkup handler...
func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
   payload := map[string]any{
		"name":"{{.PackageName}}", 
		"status-code":200, 
		"success": true, 
		"time": time.Now().Format(time.RFC3339),
	}
	

   // Return the response to the client 
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(payload)
}