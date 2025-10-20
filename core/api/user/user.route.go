package user

import (
	mw "opsie/core/middlewares"
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for user.
func HandleRoutes(r *mux.Router, h *Handler) {

	// ---------------------------------------------------------
	// PUBLIC...
	// ---------------------------------------------------------
	// Onboarding
	bolt.Api(r, bolt.MethodPost, "/owner/create", h.CreateOwnerAccount)
	bolt.Api(r, bolt.MethodGet, "/owner/count", h.GetOwnerCount)

	// Protected [Auth]
	bolt.Api(r, bolt.MethodPatch, "/account/update/name", h.UpdateAccountDisplayName, mw.Auth)
	bolt.Api(r, bolt.MethodPatch, "/account/update/password", h.UpdateAccountPassword, mw.Auth)
	bolt.Api(r, bolt.MethodPost, "/account/email/otp", h.GetOwnerCount, mw.Auth)
	bolt.Api(r, bolt.MethodPost, "/account/email/verify", h.GetOwnerCount, mw.Auth)


}
