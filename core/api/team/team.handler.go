package team

import (
	"fmt"
	"net/http"
	"opsie/core/models"
	"opsie/core/services"
	"opsie/def"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
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





/// ______________________________________________________________________________________________________
/// Protected Routes [Auth] ------------------------------------------------------------------------------
/// Accessed by all authenticated user

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
	team, err := h.service.GetUserDefaultTeam(authUser.ID)
	if err != nil {return err}



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Default teams",
		"team"			: team,
	})
	return nil
}



/// ______________________________________________________________________________________________________
/// Protected Routes [Auth, Admin] -----------------------------------------------------------------------
/// Only can be accessed by Owner & Admin 


func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Processing Request Body
	var payload models.NewTeamPayload
	bolt.ParseBody(w, r, &payload)

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




func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) *errors.Error{

	// Fetch all teams 
	teams, err := h.service.GetAll()
	if err != nil {return err}

   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All teams",
		"teams"			: teams,
	})
	return nil
}




func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) *errors.Error{

	// Get the team id form url
	id := bolt.ParseParamId(w, r, "id")


	// Fetch team by id 
	team, err1 := h.service.GetById(id)
	if err1 != nil {return err1}


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team by id",
		"team"			:  team,
	})
	return nil
}



func (h *Handler) GetAllByUserId(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the `user_id` from the URL
	userId := bolt.ParseParamId(w, r, "user_id")

	// Fetch all teams of user
	teams, err1 := h.service.GetAllByUserId(userId)
	if err1 != nil {return err1}

   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All teams of user",
		"teams"			: teams,
	})
	return nil
}



func (h *Handler) Update(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the team id form url
	id := bolt.ParseParamId(w, r, "id")

	// Processing Request Body
	var payload models.UpdateTeamPayload
	bolt.ParseBody(w, r, &payload)

	// 
	team, err1 := h.service.Update(id, payload)
	if err1 != nil {return err1}



	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team updated",
		"team"			: team,
	})
	return nil
}



func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the team id form url
	id := bolt.ParseParamId(w, r, "id")


	// Delete the team
	if e := h.service.Delete(id); e != nil {return e}


	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team Delete",
	})
	return nil
}



