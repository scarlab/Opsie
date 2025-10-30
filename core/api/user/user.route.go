package user

import (
	"opsie/core/mw"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for user.
// User Route ---
func HandleRoutes(r chi.Router, h *Handler) {
	// __________________________________________________________________________________________________
	// Onboarding [Public] ------------------------------------------------------------------------------
	api.Post(r, 			"/owner/create", 					h.CreateOwnerAccount)
	api.Get(r, 				"/owner/count", 					h.GetOwnerCount)



	// __________________________________________________________________________________________________
	// Protected [Auth] ---------------------------------------------------------------------------------
	// User Account Routes
	api.Patch(r, 			"/account/update/name", 			h.UpdateAccountDisplayName, 	mw.Auth)
	api.Patch(r, 			"/account/update/password", 		h.UpdateAccountPassword, 		mw.Auth)
	api.Post(r, 			"/account/email/otp", 				h.GetOwnerCount, 				mw.Auth) // pending...
	api.Post(r, 			"/account/email/verify", 			h.GetOwnerCount, 				mw.Auth) // pending...


	
	// __________________________________________________________________________________________________
	// Protected [Auth, Admin] --------------------------------------------------------------------------
	// Admin Route
	api.Post(r, 			"/create", 					h.Create, 						mw.Auth, mw.Admin) // pending...
	api.Get(r, 				"/get", 					h.GetAll, 						mw.Auth, mw.Admin) // pending...
	api.Get(r, 				"/get/{id}", 				h.GetByID, 						mw.Auth, mw.Admin) // pending...
	api.Patch(r, 			"/update/{id}", 			h.Update, 						mw.Auth, mw.Admin) // pending...
	api.Delete(r, 			"/delete/{id}", 			h.Delete, 						mw.Auth, mw.Admin) // pending...
	api.Post(r, 			"/team/add", 				h.AddToTeam, 					mw.Auth, mw.Admin) // pending...
	api.Post(r, 			"/team/remove", 			h.RemoveFromTeam, 				mw.Auth, mw.Admin) // pending...


}
