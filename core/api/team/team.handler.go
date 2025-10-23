package team

import (
	"fmt"
	"net/http"
	"opsie/core/models"
	"opsie/core/services"
	"opsie/def"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/logger"
)

// Team Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *services.TeamService
}

// NewHandler - Constructor for Team Handler
func NewHandler(service *services.TeamService) *Handler {
	return &Handler{
		service: service,
	}
}



func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload models.NewTeamPayload
	bolt.ParseBody(w, r, &payload)
logger.Debug("%s",payload)
	// Create Team
	team, err := h.service.Create(payload)
	if err != nil {
		return err
	}

   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team created",
		"team"	: team,
	})
	return nil
}


func (h *Handler) GetAllTeams(w http.ResponseWriter, r *http.Request) *errors.Error{


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All teams",
		"teams"		: "all",
	})
	return nil
}

func (h *Handler) GetUserTeams(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the session user
	userVal:= r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return errors.Internal(fmt.Errorf("session user not found"))
	}
	
	authUser, ok := userVal.(models.AuthUser)
	if !ok {
		return errors.Internal(fmt.Errorf("invalid session"))
	}

	// Fetch all teams of user
	teams, err := h.service.GetUserTeams(authUser.ID)
	if err != nil {return err}



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All user teams",
		"teams"		: teams,
	})
	return nil
}

func (h *Handler) GetUserDefaultTeam(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the session user
	userVal:= r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return errors.Internal(fmt.Errorf("session user not found"))
	}
	
	authUser, ok := userVal.(models.AuthUser)
	if !ok {
		return errors.Internal(fmt.Errorf("invalid session"))
	}

	// Fetch all teams of user
	teams, err := h.service.GetUserTeams(authUser.ID)
	if err != nil {return err}



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Default teams",
		"teams"		: teams,
	})
	return nil
}


func (h *Handler) UpdateInfo(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload models.UpdateTeamPayload
	bolt.ParseBody(w, r, &payload)



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team updated",
		"payload"		: payload,
	})
	return nil
}


func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload models.UpdateTeamPayload
	bolt.ParseBody(w, r, &payload)



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team deleted",
	})
	return nil
}