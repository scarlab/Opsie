package {{.PackageName}}

import (
	"opsie/pkg/api"

	"github.com/go-chi/chi/v5"
)


// HandleRoutes - Defines all HTTP endpoints for auth.
func HandleRoutes(r chi.Router, h *Handler) {
	bolt.Api(r, bolt.MethodPost, "/create", h.Create)
	// API					// Pattern							// Handler
	api.Post(r, 			"/account/email/verify", 			h.GetOwnerCount)

}

