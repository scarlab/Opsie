// internal/api/router.go
package server

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"watchtower/config"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // TODO: tighten this later
}


func (s *ApiServer) setupRouter() *mux.Router {
	router := mux.NewRouter()
	hub := NewSocketHub()


	// API routes -------------------------------------------------
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/health", healthHandler).Methods("GET")



	// Agent WebSocket endpoint
	apiRouter.HandleFunc("/ws/agent", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
			return
		}

		// ✅ First message = { "auth_key": "...", "node_id": "..." }
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}

		// TODO: validate auth_key in DB
		nodeID := string(msg) // keep simple for now
		hub.RegisterAgent(nodeID, conn)

		// Start listening in background
		go func() {
			defer func() {
				hub.UnregisterAgent(nodeID)
				conn.Close()
			}()
			for {
				_, data, err := conn.ReadMessage()
				if err != nil {
					return
				}
				// Process metrics/logs here
				fmt.Printf("From %s: %s\n", nodeID, data)

				// ✅ broadcast raw message to UI clients
				hub.BroadcastToUI(data)

				// TODO: push to DB or ...

			}
		}()
	}).Methods("GET")

	// UI WebSocket endpoint
	apiRouter.HandleFunc("/ws/ui", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
			return
		}

		hub.RegisterUI(conn)

		// Listen until closed
		go func() {
			defer func() {
				hub.UnregisterUI(conn)
				conn.Close()
			}()
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					return
				}
				// UI usually won’t send much, but you can handle commands if needed
			}
		}()
	}).Methods("GET")




	


	// Web UI Proxy
	if config.IsDev {
		viteURL, _ := url.Parse("http://localhost:5173")
		viteProxy := httputil.NewSingleHostReverseProxy(viteURL)
		router.PathPrefix("/").Handler(viteProxy)
	} else{
		// Static assets
		staticHandler := http.FileServer(http.FS(s.uiFS))
		router.PathPrefix("/assets/").Handler(staticHandler)

		// SPA fallback → index.html
		router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data, err := fs.ReadFile(s.uiFS, "index.html")
			if err != nil {
				http.Error(w, "index.html not found", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(data)
		})
	}

	return router
}




// healthHandler returns a simple health status as JSON.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}