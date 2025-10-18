// ================================================================
// Author: Samrat
// Created: 2025/09/21 10:57:53
// ================================================================

package ws_agent

import (
	"database/sql"
	"opsie/internal/socket"

	"github.com/gorilla/mux"
)

// Init - Entry point for initializing socket connection with agent
func Register(r *mux.Router, db *sql.DB, socketHub *socket.Hub) {

	// Agent
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service, socketHub)
	r.HandleFunc("/agent", handler.Connect).Methods("GET")
}
