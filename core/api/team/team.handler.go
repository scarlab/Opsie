package team

import (
	"net/http"
	"opsie/core/models"
	"opsie/core/repo"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
)

// Team Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	repo *repo.TeamRepository
	userTeamRepo *repo.UserTeamRepository
}

// NewHandler - Constructor for Team Handler
func NewHandler(
	repo *repo.TeamRepository, 
	userTeamRepo *repo.UserTeamRepository,
	) *Handler {
	return &Handler{
		repo: repo,
		userTeamRepo: userTeamRepo,
	}
}





/// ______________________________________________________________________________________________________
/// Protected Routes [Auth] ------------------------------------------------------------------------------
/// Accessed by all authenticated user

func (h *Handler) GetUserTeams(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the request/session user 
	sessionUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	// Fetch all teams of user
	teams, err := h.userTeamRepo.ListTeamsByUser(sessionUser.ID)
	if err != nil {return err}


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "All user teams",
		"teams"		: teams,
	})
	return nil
}




func (h *Handler) GetUserDefaultTeam(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the request/session user 
	sessionUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	// Fetch all teams of user
	team, err := h.userTeamRepo.DefaultTeam(sessionUser.ID)
	if err != nil {return err}




   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Default teams",
		"team"			: team,
	})
	return nil
}

func (h *Handler) SetUserDefaultTeam(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the request/session user 
	sessionUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	id := bolt.ParseParamId(w, r, "id")

	// Fetch all teams of user
	team, err := h.userTeamRepo.SetDefaultTeam(sessionUser.ID, id)
	if err != nil {return err}



   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team switched",
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
	team, err := h.repo.Create(payload)
	if err != nil { return err }


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team created",
		"team"	: team,
	})
	return nil
}




func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) *errors.Error{

	// Fetch all teams 
	teams, err := h.repo.GetAll()
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
	team, err1 := h.repo.GetById(id)
	if err1 != nil {return err1}


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team by id",
		"team"			:  team,
	})
	return nil
}


func (h *Handler) GetAllMembersOfTeam(w http.ResponseWriter, r *http.Request) *errors.Error{

	// Get the team id form url
	team_id := bolt.ParseParamId(w, r, "team_id")


	// Fetch team by id 
	members, err1 := h.userTeamRepo.ListTeamMembers(team_id)
	if err1 != nil {return err1}


   	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team members",
		"members"		:  members,
	})
	return nil
}



func (h *Handler) GetAllByUserId(w http.ResponseWriter, r *http.Request) *errors.Error{
	// Get the `user_id` from the URL
	userId := bolt.ParseParamId(w, r, "user_id")

	// Fetch all teams of user
	teams, err1 := h.repo.GetAllByUserId(userId)
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
	var payload models.NewTeamPayload
	bolt.ParseBody(w, r, &payload)

	// 
	team, err1 := h.repo.Update(id, payload)
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
	if e := h.repo.Delete(id); e != nil {return e}


	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message"		: "Team Delete",
	})
	return nil
}



