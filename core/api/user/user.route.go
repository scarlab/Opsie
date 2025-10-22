package user

import (
	"opsie/core/mw"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for user.
func HandleRoutes(r chi.Router, h *Handler) {

	// ---------------------------------------------------------
	// PUBLIC...
	// ---------------------------------------------------------
	// Onboarding
	api.Post(r, 			"/owner/create", 					h.CreateOwnerAccount)
	api.Get(r, 				"/owner/count", 					h.GetOwnerCount)

	// Protected [Auth]
	api.Patch(r, 			"/account/update/name", 			h.UpdateAccountDisplayName, 	mw.Auth)
	api.Patch(r, 			"/account/update/password", 		h.UpdateAccountPassword, 		mw.Auth)
	api.Post(r, 			"/account/email/otp", 				h.GetOwnerCount, 				mw.Auth)
	api.Post(r, 			"/account/email/verify", 			h.GetOwnerCount, 				mw.Auth)


}
