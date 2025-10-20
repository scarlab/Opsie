package organization

import (
	"net/http"
	"opsie/core/services"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/logger"
	"opsie/types"
)

// Organization Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *services.OrganizationService
}

// NewHandler - Constructor for Organization Handler
func NewHandler(service *services.OrganizationService) *Handler {
	return &Handler{
		service: service,
	}
}



func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload types.NewOrganizationPayload
	bolt.ParseBody(w, r, &payload)
logger.Debug("%s",payload)
	// Create Organization
	organization, err := h.service.Create(payload)
	if err != nil {
		return err
	}

   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Organization created",
		"organization"	: organization,
	})
	return nil
}