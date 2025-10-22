package team

import (
	"opsie/core/mw"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r chi.Router, h *Handler) {
	// Stuff [Auth] 
	api.Get(r,  			"/get/user/all", 				h.GetUserTeams, 			mw.Auth)
	api.Get(r,  			"/get/user/default", 			h.GetUserDefaultTeam, 		mw.Auth)

	// Owner/Admin [Auth, Admin]
	api.Post(r,  			"/create", 						h.Create,							mw.Auth)
	api.Patch(r,  			"/updated/info", 				h.UpdateInfo,						mw.Auth)
	api.Get(r,  			"/get", 						h.GetAllTeams, 				mw.Auth)
	api.Delete(r,  			"/delete", 						h.Delete,							mw.Auth)
}

