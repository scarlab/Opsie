package organization

import (
	mw "opsie/core/middlewares"
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)

// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r chi.Router, h *Handler) {
	api.Post(r,  	"/create", 			h.Create,					mw.Auth)
	api.Patch(r,  	"/updated/info", 	h.UpdateInfo,				mw.Auth)
	api.Get(r,  	"/get", 			h.GetAllOrganizations, 		mw.Auth)
	api.Get(r,  	"/get/user-orgs", 	h.GetUserOrganizations, 	mw.Auth)
	api.Delete(r,  	"/delete", 			h.Delete,					mw.Auth)
}

