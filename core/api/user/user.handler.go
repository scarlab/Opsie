package user

import (
	"net/http"
	"opsie/config"
	"opsie/core/models"
	"opsie/core/services"
	"opsie/def"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
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


/// ______________________________________________________________________________________________________
/// Public Routes ----------------------------------------------------------------------------------------
/// --- 

func (h *Handler) CreateOwnerAccount(w http.ResponseWriter, r *http.Request) *errors.Error {
    // Processing Request Body
	var payload models.NewUserPayload
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

func (h *Handler) GetOwnerCount(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Handling Business Logics
	count, err := h.service.GetOwnerCount()
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
	authUser, err := h.service.UpdateAccountName(sessionUser.ID, payload)
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
	sessionUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	// Get the session
	session, gsErr := bolt.GetSession(r)
	if gsErr!= nil {return gsErr}

	// Processing Request Body
	var payload models.UpdateAccountPasswordPayload
	bolt.ParseBody(w, r, &payload)
	
	// Handling Business Logics
	newSession, err := h.service.UpdateAccountPassword(sessionUser.ID, session.Key, payload)
	if err != nil { return err}


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

	user, err := h.service.CreateUser(payload)
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
	users, err := h.service.GetAllUser()
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
	user, err := h.service.GetUserById(id)
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
	user, err := h.service.GetUserById(id)
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
	if e := h.service.DeleteUser(id); e != nil {return e}
	
	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User deleted!",
		"deleted": true,
	})

	return nil
}


func (h *Handler) AddToTeam(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing id from URL params
	userId := bolt.ParseParamId(w, r, "user_id")
	teamId := bolt.ParseParamId(w, r, "team_id")


	// Create team User
	if e := h.service.AddUserToTeam(userId, teamId); e != nil {return e}

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message"		: "User added to the team",
	})

	return nil
}


func (h *Handler) RemoveFromTeam(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Processing id from URL params
	userId := bolt.ParseParamId(w, r, "user_id")
	teamId := bolt.ParseParamId(w, r, "team_id")


	// Delete UserTeam record
	if e := h.service.AddUserToTeam(userId, teamId); e != nil {return e}
	

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "User deleted!",
		"all_user": nil,
	})

	return nil
}