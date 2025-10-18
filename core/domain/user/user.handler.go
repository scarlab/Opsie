package user

import (
	"net/http"
	"opsie/core/services"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/types"
)

// Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *services.UserService
}
 
// NewHandler - Constructor for Handler
func NewHandler(service *services.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateOwnerAccount(w http.ResponseWriter, r *http.Request) *errors.Error {
    // Processing Request Body
	var payload types.NewOwnerPayload
	bolt.ParseBody(w, r, &payload)

	// Handling Business Logics
	user, err := h.service.CreateOwnerAccount(payload)
	if err != nil { return err}


	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "Owner account created!",
		"user":    user,
	})

	return nil
}



