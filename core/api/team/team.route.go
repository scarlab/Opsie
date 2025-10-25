package team

import (
	"opsie/core/mw"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
// Team Route ---
func HandleRoutes(r chi.Router, h *Handler) {
	/// _________________________________________________________________________________________________
	/// Protected [Auth] --------------------------------------------------------------------------------
	/// For Staff
	api.Get(r,  			"/get/user/all", 				h.GetUserTeams, 			mw.Auth)
	api.Get(r,  			"/get/user/default", 			h.GetUserDefaultTeam, 		mw.Auth)



	/// _________________________________________________________________________________________________
	/// Protected [Auth, Admin] -------------------------------------------------------------------------
	/// Team maintenance
	api.Post(r,  			"/create", 						h.Create,					mw.Auth, mw.Admin) 
	api.Get(r,  			"/get", 						h.GetAll, 					mw.Auth, mw.Admin)
	api.Get(r,  			"/get/{id}", 					h.GetById, 					mw.Auth, mw.Admin)
	api.Get(r,  			"/get/user/{user_id}", 			h.GetAllByUserId, 			mw.Auth, mw.Admin)
	api.Patch(r,  			"/update/{id}", 				h.Update,					mw.Auth, mw.Admin)
	api.Delete(r,  			"/delete/{id}", 				h.Delete,					mw.Auth, mw.Admin) // pending...
}

 