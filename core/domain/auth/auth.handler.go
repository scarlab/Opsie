package auth

import (
	"fmt"
	"net/http"
	"opsie/config"
	"opsie/constant"
	"opsie/core/services"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/types"
)

// Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *services.AuthService
}

// NewHandler - Constructor for Handler
func NewHandler(service *services.AuthService) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) Login(w http.ResponseWriter, r *http.Request) *errors.Error{
    // Processing Request Body
	var payload types.LoginPayload
	bolt.ParseBody(w, r, &payload)

	
	// Handling Business Logics
	authUser, err := h.service.AuthenticateUser(payload)
	if err != nil { return err }


	// Create Session
	session, err := h.service.CreateSession(authUser.ID)
	if err != nil { return err }
	

	// Set Headers/Cookies
	// set cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session",
        Value:    session.Key,
        Expires:  session.Expiry,
        HttpOnly: true,
        Secure:   !config.IsDev,
        Path:     "/",
        SameSite: http.SameSiteLaxMode,
    })

	// Send the final response 
	bolt.WriteResponse(w, 200, map[string]any{
		"message": "Login Successful",
		"auth_user":    authUser,
		"session_key":    session.Key,
	})
	return nil
}


func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) *errors.Error{
    

	fmt.Println("user - ", r.Context().Value(constant.ContextKeyUser))
	fmt.Println("session - ", r.Context().Value(constant.ContextKeySession))
	// Send the final response 
	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message": "Successfully logged out",
	})
	return nil
}
