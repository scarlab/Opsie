package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"opsie/config"
	"opsie/internal/socket"
	ws_agent "opsie/internal/socket/clients/agent"
	ws_ui "opsie/internal/socket/clients/ui"

	"github.com/gorilla/mux"
)





func (s *ApiServer) setupRouter() *mux.Router {
	router := mux.NewRouter()
	socketHub := socket.NewHub()

	// ------------------------------------ Root API routes ------------------------------------ //
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/health", healthHandler).Methods("GET")


	// ----------------------------------- Web Socket Routes ----------------------------------- //
	wsRouter := apiRouter.PathPrefix("/ws").Subrouter()
	ws_agent.Init(wsRouter, s.db, socketHub)
	ws_ui.Init(wsRouter, s.db, socketHub)


	// ---------------------------------------- Domains ---------------------------------------- //
	// ws.Init(apiRouter, s.db, socketHub) // Web Socket Communication



	// ------------------------------------- Web UI Proxy ------------------------------------- //
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