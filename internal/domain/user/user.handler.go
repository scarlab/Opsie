package user

import (
	"encoding/json"
	"net/http"
	"opsie/pkg/bolt"

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

func (h *Handler) CreateOwnerAccount(w http.ResponseWriter, r *http.Request) {
    // Processing Request Body
	var payload TNewOwnerPayload
	bolt.ParseBody(w, r, &payload)

	// Handling Business Logics
	user, err := h.service.CreateOwnerAccount(payload)
	if bolt.ErrorHandler(w, err) { return }


	// Send the final response 
	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message": "Owner account created!",
		"user":    user,
	})
}



// Health checkup handler...
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
   payload := map[string]any{
		"name":"user", 
		"status-code":200, 
		"success": true, 
		"time": time.Now().Format(time.RFC3339),
	}
	

   // Return the response to the client 
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(payload)
}