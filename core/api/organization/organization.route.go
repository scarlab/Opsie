package organization

import (
	mw "opsie/core/middlewares"
	"opsie/pkg/bolt"

	"github.com/gorilla/mux"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r *mux.Router, h *Handler) {
	bolt.Api(r, bolt.MethodPost, 	"/create", 			h.Create,					mw.Auth)
	bolt.Api(r, bolt.MethodPatch, 	"/updated/info", 	h.UpdateInfo,				mw.Auth)
	bolt.Api(r, bolt.MethodGet, 	"/get", 			h.GetAllOrganizations, 		mw.Auth)
	bolt.Api(r, bolt.MethodGet, 	"/get/user-orgs", 	h.GetUserOrganizations, 	mw.Auth)
	bolt.Api(r, bolt.MethodDelete, 	"/delete", 			h.Delete,					mw.Auth)
}

