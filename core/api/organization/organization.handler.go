package organization

import (
	"fmt"
	"net/http"
	"opsie/core/services"
	"opsie/def"
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


func (h *Handler) GetAllOrganizations(w http.ResponseWriter, r *http.Request) *errors.Error{


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All organizations",
		"organizations"		: "all",
	})
	return nil
}

func (h *Handler) GetUserOrganizations(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the session user
	userVal:= r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return errors.Internal(fmt.Errorf("session user not found"))
	}
	
	authUser, ok := userVal.(types.AuthUser)
	if !ok {
		return errors.Internal(fmt.Errorf("invalid session"))
	}

	// Fetch all orgs of user
	orgs, err := h.service.GetUserOrganizations(authUser.ID)
	if err != nil {return err}



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All user organizations",
		"organizations"		: orgs,
	})
	return nil
}

func (h *Handler) GetUserDefaultOrganization(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the session user
	userVal:= r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return errors.Internal(fmt.Errorf("session user not found"))
	}
	
	authUser, ok := userVal.(types.AuthUser)
	if !ok {
		return errors.Internal(fmt.Errorf("invalid session"))
	}

	// Fetch all orgs of user
	orgs, err := h.service.GetUserOrganizations(authUser.ID)
	if err != nil {return err}



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Default organizations",
		"organizations"		: orgs,
	})
	return nil
}


func (h *Handler) UpdateInfo(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload types.UpdateOrganizationPayload
	bolt.ParseBody(w, r, &payload)



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Organization updated",
		"payload"		: payload,
	})
	return nil
}


func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload types.UpdateOrganizationPayload
	bolt.ParseBody(w, r, &payload)



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Organization deleted",
	})
	return nil
}