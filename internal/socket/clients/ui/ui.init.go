// ================================================================
// Author: Samrat
// Created: 2025/09/21 11:25:57
// ================================================================

package ws_ui

import (
	"database/sql"
	"watchtower/internal/socket"

	"github.com/gorilla/mux"
)

// Init - Entry point for initializing socket connection with agent
func Init(r *mux.Router, db *sql.DB, socketHub *socket.Hub) {

	// Agent
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service, socketHub)
	r.HandleFunc("/ui", handler.Connect).Methods("GET")
}
