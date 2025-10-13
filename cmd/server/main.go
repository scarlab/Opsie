package main

import (
	"context"
	"io/fs"
	"log"
	"net"
	embedui "opsie"
	"opsie/config"
	"opsie/db"
	"os"
	"os/signal"
	"syscall"

	"opsie/internal/server"
	"opsie/internal/socket"
)

// getLocalIP returns the first non-loopback local IP address (e.g., 192.168.x.x).
// If no network interface is found, it falls back to "localhost".
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}


// main is the entry point of the Clann backend server.
// It initializes all core services (database, WebSocket hub, API server)
// and gracefully handles shutdown signals.
func main() {
	log.Printf("🚀 Starting %s (%s)\n", config.AppConfig.Name, config.AppConfig.Version)
	log.Printf("🌱 Environment: %s\n\n", config.ENV.GoEnv)


	// -------------------------------------------------------------------
	// Embed React UI
	// -------------------------------------------------------------------
	uiFS, err := fs.Sub(embedui.EmbeddedUI, "ui/dist")
	if err != nil {
		log.Fatalf("💀 Web UI Embedding failed: %v", err)
	}
	log.Printf("✅ Web UI ready")
	
	// -------------------------------------------------------------------
	// Initialize Database
	// -------------------------------------------------------------------
	database, err := db.Postgres()
	if err != nil {
		log.Fatalf("💀 Database connection failed: %v", err)
	}

	
	// -------------------------------------------------------------------
	// Initialize WebSocket Hub
	// -------------------------------------------------------------------
	socketHub := socket.NewHub()
	log.Println("✅ WebSocket hub ready")


	// -------------------------------------------------------------------
	// Setup Context & Signal Handling (graceful shutdown)
	// -------------------------------------------------------------------
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()



	// -------------------------------------------------------------------
	// Start API Server
	// -------------------------------------------------------------------
	apiServer := server.NewApiServer(config.ENV.Addr, database,uiFS, socketHub)

	log.Printf("🌍 Server listening on http://%s%s\n", GetLocalIP(),config.ENV.Addr)


	// Blocking call — will run until context is cancelled
	if err := apiServer.Run(ctx); err != nil {
		log.Fatalf("💀 Server encountered an error: %v", err)
	}

	log.Println("🛑 Server shutdown complete. Goodbye 👋")
}
