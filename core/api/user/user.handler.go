package user

import (
	"net/http"
	"opsie/config"
	"opsie/core/services"
	"opsie/def"
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


// --------------------------------------------------------------------------------
// Public Routes
// --------------------------------------------------------------------------------

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


// --------------------------------------------------------------------------------
// Protected Routes [Auth] 
// --------------------------------------------------------------------------------

func (h *Handler) UpdateAccountDisplayName(w http.ResponseWriter, r *http.Request) *errors.Error {
	// Get the request/session user
	sessionUser, gsuErr := bolt.GetSessionUser(r)
	if gsuErr!= nil {return gsuErr}

	// Processing Request Body
	var payload types.UpdateAccountNamePayload
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
	var payload types.UpdateAccountPasswordPayload
	bolt.ParseBody(w, r, &payload)
	
	// Handling Business Logics
	newSession, err := h.service.UpdateAccountPassword(sessionUser.ID, session.Key, payload)
	if err != nil { return err}


	// Set Headers/Cookies
	// set cookie
    http.SetCookie(w, &http.Cookie{
        Name:     def.CookieNameSession,
        Value:    newSession.Key.ToString(),
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