package auth

import (
	"net/http"
	"opsie/config"
	"opsie/pkg/bolt"
)

// Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	service *Service
}

// NewHandler - Constructor for Handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    // Processing Request Body
	var payload TLoginPayload
	bolt.ParseBody(w, r, &payload)

	
	// Handling Business Logics
	authUser, err := h.service.AuthenticateUser(payload)
	if bolt.ErrorHandler(w, err) { return }

	
	// Create Session
	session, err := h.service.CreateSession(authUser.ID)
	if bolt.ErrorHandler(w, err) { return }
	

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
}


func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
    // Processing Request Body
	// var payload TNewOwnerPayload
	// bolt.ParseBody(w, r, &payload)

	// Handling Business Logics
	// _, err := h.service.CreateOwnerAccount(payload)
	// if bolt.ErrorHandler(w, err) { return }


	// Send the final response 
	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message": "Successfully logged out",
	})
}
