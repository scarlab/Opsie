package user

import (
	"net/http"
	"opsie/config"
	"opsie/core/models"
	"opsie/core/repo"
	"opsie/def"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
)

// Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	repo *repo.UserRepository
	authRepo *repo.AuthRepository
	teamRepo *repo.TeamRepository
	userTeamRepo *repo.UserTeamRepository
}
 
// NewHandler - Constructor for Handler
func NewHandler(
	repo *repo.UserRepository, 
	authRepo *repo.AuthRepository,
	teamRepo *repo.TeamRepository, 
	userTeamRepo *repo.UserTeamRepository,
	) *Handler {
	return &Handler{
		repo: repo,
		authRepo: authRepo,
		teamRepo: teamRepo,
		userTeamRepo: userTeamRepo,
	}
}


/// ______________________________________________________________________________________________________
/// Public Routes ----------------------------------------------------------------------------------------
/// --- 

func (h *Handler) CreateOwnerAccount(w http.ResponseWriter, r *http.Request) *errors.Error { 
    // Processing Request Body
	var payload models.NewUserPayload
	bolt.ParseBody(w, r, &payload)

	// validate payload
	if payload.Email == "" || payload.Password == "" {
		return  errors.BadRequest("Email and password required")
	}

	// Generate the hashed password
	hashedPassword, _ 	:= utils.Hash.Generate(payload.Password)


	// Update Owner Payload
	payload.Password 	= hashedPassword
	payload.SystemRole 	= def.SystemRoleOwner.ToString()
	payload.ResetPass	= false


	// Create the Owner account
	user, err := h.repo.Create(payload)
	if err != nil { return err}


	// Create default team of owner
	teamPayload := models.NewTeamPayload{
		Name:        utils.GenerateTeamName(),
		Description: "This is your default team.",
	}
	team, teamErr := h.teamRepo.Create( teamPayload)
	if teamErr != nil {
		return  teamErr
	}

	userTeamPayload := models.AddUserToTeamPayload{
		UserID: user.ID,
		TeamID: team.ID,
		IsDefault: true,
		IsAdmin: true,
		InvitedBy: nil,
	}

	// Link user <-> team
	if addErr := h.userTeamRepo.AddUserToTeam(userTeamPayload); addErr != nil {
		return  addErr
	}



	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "Owner account created!",
		"user":    user,
	})

	return nil
}

func (h *Handler) GetOwnerCount(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Handling Business Logics
	count, err := h.repo.GetOwnerCount()
	if err != nil { return err}

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "Owner count",
		"count":    count,
	})

	return nil
}


/// ______________________________________________________________________________________________________
/// Protected Routes [Auth] ------------------------------------------------------------------------------
/// User Account: Every user can access. 

func (h *Handler) UpdateAccountDisplayName(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Get the request/session user 
	sessionUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	// Processing Request Body
	var payload models.UpdateAccountNamePayload
	bolt.ParseBody(w, r, &payload)


	// Handling Business Logics
	authUser, err := h.repo.UpdateAccountName(sessionUser.ID, payload.DisplayName)
	if err != nil { return err}

	
	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "Name Updated",
		"auth_user":    authUser,
	})

	return nil
}

func (h *Handler) UpdateAccountPassword(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Get the request/session user
	reqUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	// Get the session
	reqSession, gsErr := bolt.GetSession(r)
	if gsErr!= nil {return gsErr}

	// Processing Request Body
	var payload models.UpdateAccountPasswordPayload
	bolt.ParseBody(w, r, &payload)
	
	// Handling Business Logics
	_, err := h.repo.UpdateAccountPassword(reqUser.ID, payload.NewPassword)
	if err != nil { return err}


	// Regenerate Session key
	newSession, rskErr := h.authRepo.RegenerateSessionKey(reqSession.Key)
	if rskErr != nil {
		return err
	}

	// Set Headers/Cookies
	// set cookie
    http.SetCookie(w, &http.Cookie{
        Name:     def.CookieNameSession,
        Value:    newSession.Key,
        Expires:  newSession.Expiry,
        HttpOnly: true,
        Secure:   !config.IsDev,
        Path:     "/",
        SameSite: http.SameSiteLaxMode,
    })

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "Password Updated",
		"session_key":    newSession.Key,
	})

	return nil
}



/// ______________________________________________________________________________________________________
/// Protected Routes [Auth, Admin] -----------------------------------------------------------------------
/// User Maintenance: Only can be accessed by Owner & Admin 


func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing Request Body
	var payload models.NewUserPayload
	bolt.ParseBody(w, r, &payload)

	// validate payload
	if payload.Email == "" || payload.Password == "" {
		return  errors.BadRequest("Email and password required")
	}

	// Generate the hashed password
	hashedPassword, _ 	:= utils.Hash.Generate(payload.Password)
	payload.Password = hashedPassword
	

	user, err := h.repo.Create(payload)
	if err != nil {return err}

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User Created",
		"user": user,
	})

	return nil
}


func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Get all users
	users, err := h.repo.GetAll()
	if err != nil {return err}


	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "List of all users",
		"users": users,
	})

	return nil
}


func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing id from URL params
	id := bolt.ParseParamId(w, r, "id")

	// Get the user with ID
	user, err := h.repo.GetByID(id)
	if err != nil {return err}

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User by id",
		"user": user,
	})

	return nil
}


func (h *Handler) Update(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing Request Body
	var payload models.UpdateUserPayload
	bolt.ParseBody(w, r, &payload)

	// Processing id from URL params
	id := bolt.ParseParamId(w, r, "id")

	// Update the user with ID
	user, err := h.repo.GetByID(id)
	if err != nil {return err}

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message"		: "User updated!",
		"user"			: user,
	})

	return nil
}


func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing id from URL params
	id := bolt.ParseParamId(w, r, "id")

	// Delete the team
	if e := h.repo.Delete(id); e != nil {return e}
	
	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User deleted!",
		"deleted": true,
	})

	return nil
}


func (h *Handler) AddToTeam(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing ids from URL params
	var payload models.AddUserToTeamPayload
	bolt.ParseBody(w, r, &payload)
	

	// Create team User
	if e := h.userTeamRepo.AddUserToTeam(payload); e != nil {return e}

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message"		: "User added to the team",
	})

	return nil
}


func (h *Handler) RemoveFromTeam(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing ids from URL params
	userId := bolt.ParseParamId(w, r, "user_id")
	teamId := bolt.ParseParamId(w, r, "team_id")


	// Delete UserTeam record
	if e := h.userTeamRepo.RemoveUserFromTeam(userId, teamId); e != nil {return e}
	

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User deleted!",
		"all_user": nil,
	})

	return nil
}



func (h *Handler) RemoveAllUserFromTeam(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing id from URL params
	teamId := bolt.ParseParamId(w, r, "team_id")


	// Delete UserTeam record
	if e := h.userTeamRepo.RemoveAllUserFromTeam(teamId); e != nil {return e}
	

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User deleted!",
		"all_user": nil,
	})

	return nil
}