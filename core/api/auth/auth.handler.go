package auth

import (
	"fmt"
	"net/http"
	"opsie/config"
	"opsie/core/models"
	"opsie/core/repo"
	"opsie/def"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"time"
)

// Handler - Handles HTTP requests & responses.
// Talks only to the Service layer, not directly to Repository.
type Handler struct {
	repo *repo.AuthRepository
	userRepo *repo.UserRepository
}

// NewHandler - Constructor for Handler
func NewHandler(repo *repo.AuthRepository, userRepo *repo.UserRepository) *Handler {
	return &Handler{
		repo: repo,
		userRepo: userRepo,
	}
}


func (h *Handler) Login(w http.ResponseWriter, r *http.Request) *errors.Error{
    // Processing Request Body
	var payload models.LoginPayload
	bolt.ParseBody(w, r, &payload)

	
	// Handling Business Logics
	// authUser, err := h.service.AuthenticateUser(payload)
	// Get Request User By Email
	reqUser, err := h.userRepo.GetByEmail(payload.Email)
	if err != nil {
		return err
	}

	// Compare Password
	isMatched := utils.Hash.Compare(reqUser.Password, payload.Password)
	
	if !isMatched {
	  return  errors.Unauthorized("invalid email or password")
	}


	// Create Session
	// session, err := h.service.CreateSession(authUser.ID)
	key, sKeyErr := utils.GenerateSessionKey()
	if sKeyErr != nil {	errors.Internal(sKeyErr) }

	expiry := time.Now().Add(time.Duration(config.App.SessionDays) * 24 * time.Hour)

	session, err1 := h.repo.CreateSession(reqUser.ID, key, expiry)
	if err1 != nil { return err1 }
	

	// Generate Auth user
	authUser := models.AuthUser{
		ID: reqUser.ID,
		DisplayName: reqUser.DisplayName,
		Email: reqUser.Email,
		Avatar: reqUser.Avatar,
		SystemRole: reqUser.SystemRole,
		Preference: reqUser.Preference,
	}


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


func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) *errors.Error {
	// 1. Get session from context
	sessionVal := r.Context().Value(def.ContextKeySession)
	if sessionVal == nil {
		return errors.New(http.StatusUnauthorized, "no active session found")
	}
	session, ok := sessionVal.(models.Session)
	if !ok {
		return errors.Internal(fmt.Errorf("invalid session context type"))
	}

	// 2. Call service to invalidate the session
	if err := h.repo.ExpireSession(session.Key); err != nil {
		return err
	}

	// 3. Delete the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1, // Delete immediately
	})

	// 4. Respond success
	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message": "Successfully logged out",
	})

	return nil
}



func (h *Handler) GetSessionUser(w http.ResponseWriter, r *http.Request) *errors.Error {
	// 1. Get session from context
	userVal := r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return errors.Internal(fmt.Errorf("session user not found"))
	}

	
	authUser, ok := userVal.(models.AuthUser)
	if !ok {
		return errors.Internal(fmt.Errorf("invalid session"))
	}
	
	// time.Sleep(2 * time.Second)
	// 4. Respond success
	bolt.WriteResponse(w, http.StatusOK, map[string]any{
		"message": "Authenticated User",
		"auth_user": authUser,
	})

	return nil
}

