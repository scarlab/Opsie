package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
)

// ApiServer represents the main HTTP server for the application.
// It contains the listening address and a reference to the database connection.
type ApiServer struct {
	addr string   // Address where the server will listen (e.g. ":8080")
	db   *sql.DB  // Database connection pool
}


// Constructor: NewApiServer creates and returns a new ApiServer instance.
func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}


// Run starts the HTTP server and listens for incoming requests.
// It also listens for context cancellation to gracefully shut down the server.
func (s *ApiServer) Run(ctx context.Context) error {

	router := s.setupRouter()

	// Configure HTTP server
	server := &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	// Run the server in a separate goroutine
	go func() {
		log.Printf("✔️  Server is listening on http://localhost%s", s.addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[CRASH]: %v", err)
		}
	}()

	// Wait for context cancellation (e.g. SIGTERM or CTRL+C)
	<-ctx.Done()
	log.Println("Shutdown signal received...")

	// Gracefully shutdown the server with a timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Failed to shutdown server gracefully: %v", err)
	}

	log.Println("✔️  Server stopped gracefully")
	return nil
}


