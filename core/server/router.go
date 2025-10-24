package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"opsie/config"
	"opsie/core/api/auth"
	"opsie/core/api/team"
	"opsie/core/api/user"
	"opsie/core/mw"
	"opsie/pkg/api"
	"opsie/pkg/errors"
	"opsie/pkg/utils"

	"github.com/go-chi/chi/v5"
)



func (s *ApiServer) Router() chi.Router {
	// ___________________________________________________________________
	// Root Router -------------------------------------------------------
	/// --- 
	var router = chi.NewRouter()


	
	/// ___________________________________________________________________
	/// Web UI ------------------------------------------------------------
	/// Dev : Proxy vite app
	/// Prod: Embed vite build app
	setupUI(router, s.uiFS)



	/// ___________________________________________________________________
	/// File Server -------------------------------------------------------
	/// Serves static files for the `static_dir` dir
	/// Dev : "./uploads"
	/// Prod: "/var/lib/opsie/static"
	setupFileServer(router)



	/// ___________________________________________________________________
	/// Middlewares -------------------------------------------------------
	/// --- 
	mw.Register(s.db)



	/// ___________________________________________________________________
	/// API Gateway [v1] --------------------------------------------------
	/// --- 
	router.Route("/api/v1", func(apiRoute chi.Router) {
		// Api Home Root
		api.Get(apiRoute, "", apiHome)



		/// API: ___________________________________________________________
		
		// Authentication Routes 
		apiRoute.Route("/auth",func(r chi.Router){
			auth.Register(r, s.db)
		})
		
		// User Routes
		apiRoute.Route("/user",func(r chi.Router){
			user.Register(r, s.db)
		})
		
		// Team Routes
		apiRoute.Route("/team",func(r chi.Router){
			team.Register(r, s.db)
		})


		/// ___________________________________________________________________
		/// Handle Unknown API endpoint: 404! ---------------------------------
		/// --- 
		apiRoute.NotFound(notFound)
	})

	return router
}




// healthHandler returns a simple apiHome status as JSON.
func apiHome(w http.ResponseWriter, r *http.Request) *errors.Error{
	msg := map[string]interface{}{"status": "ok", "id":utils.GenerateID()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

	return nil
}


// Not Found: 404 Handler
func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	resp := map[string]interface{}{
		"error": "API endpoint not found",
		"code":    http.StatusNotFound,
		"method":  r.Method,
		"path":    r.URL.Path,
	}

	_ = json.NewEncoder(w).Encode(resp)
}





// Proxy/Embed UI
func setupUI(router chi.Router, uiFS fs.FS) {
	if config.IsDev {
		// Dev: Proxy to Vite server
		viteURL, _ := url.Parse("http://" + config.ENV.DevUIHost + ":5173/")
		viteProxy := httputil.NewSingleHostReverseProxy(viteURL)
		router.Handle("/*", viteProxy)
	} else if config.IsProd {
		// Prod: Serve static assets
		router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(uiFS))))

		// SPA fallback → serve index.html for all other routes
		router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
			data, err := fs.ReadFile(uiFS, "index.html")
			if err != nil {
				http.Error(w, "index.html not found", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(data)
		})
	}
}

func setupFileServer(router chi.Router) {
	// Serve _static directory with your middleware
	fileServer := http.FileServer(http.Dir(config.DefaultStaticDir))
	staticHandler := http.StripPrefix("/_static/", fileServer)

	// Wrap with your bolt middleware
	// mwh := bolt.AdaptToHTTP(bolt.HandleMiddleware(bolt.AdaptFromHTTP(staticHandler)))

	router.Handle("/_static/*", staticHandler) // Chi uses Handle with wildcard
}