// internal/api/router.go
package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *ApiServer) setupRouter() *mux.Router {
	// Create a new router
	router := mux.NewRouter()

	// Versioned API sub-router or base-router
	baseRouter := router.PathPrefix("/api/v1").Subrouter()

	// Health check endpoint
	baseRouter.HandleFunc("/health", healthHandler).Methods("GET")

	// Register domain routes ------------------------------------------------
	// auth.Init(baseRouter, s.db)
	

	return router
}



// healthHandler returns a simple health status as JSON.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}