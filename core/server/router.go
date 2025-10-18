package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"opsie/config"
	"opsie/internal/domain/auth"
	"opsie/internal/domain/user"
	ws_agent "opsie/internal/socket/clients/agent"
	ws_ui "opsie/internal/socket/clients/ui"
	"opsie/pkg/bolt"
	"opsie/pkg/errors"
	"opsie/pkg/utils"

	"github.com/gorilla/mux"
)






func (s *ApiServer) Router() *mux.Router {
	// -------------------------------------------------------------------
	// Root Router
	// -------------------------------------------------------------------
	var router = mux.NewRouter()


	// -------------------------------------------------------------------
	// Gateway of API routes
	// -------------------------------------------------------------------
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	bolt.Api(apiRouter, "GET", "", apiHome)
	

	// -------------------------------------------------------------------
	// Web Socket Routes
	// -------------------------------------------------------------------
	wsRouter := apiRouter.PathPrefix("/ws").Subrouter()
	ws_agent.Register(wsRouter, s.db, s.socketHub)
	ws_ui.Register(wsRouter, s.db, s.socketHub)


	// -------------------------------------------------------------------
	// Register Domains
	// -------------------------------------------------------------------
	user.Register(apiRouter, s.db)
	auth.Register(apiRouter, s.db)


	// -------------------------------------------------------------------
	// Handle Unknown API endpoint 404
	// -------------------------------------------------------------------
	router.PathPrefix("/api/").HandlerFunc(bolt.NormalizedMiddleware(bolt.Middleware(notFound)))



	// -------------------------------------------------------------------
	// Web UI - Proxy(dev) / Embed (prod)
	// -------------------------------------------------------------------
	if config.IsDev {
		viteURL, _ := url.Parse("http://"+config.ENV.DevUIHost+":5173/")
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


// healthHandler returns a simple apiHome status as JSON.
func apiHome(w http.ResponseWriter, r *http.Request) *errors.Error{
	msg := map[string]interface{}{"status": "ok", "id":utils.GenerateID()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

	return nil
}

func notFound(w http.ResponseWriter, r *http.Request) *errors.Error{
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	resp := map[string]interface{}{
		"error": "API endpoint not found",
		"code":    http.StatusNotFound,
		"path":    r.URL.Path,
		"method":  r.Method,
	}

	_ = json.NewEncoder(w).Encode(resp)

	return nil
}
