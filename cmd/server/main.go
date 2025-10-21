package main

import (
	"context"
	"io/fs"
	"log"
	"net"
	embedui "opsie"
	"opsie/config"
	"opsie/db"
	"opsie/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"opsie/core/server"
	"opsie/core/socket"
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
	logger.Init()

	logger.Info("🚀 Starting %s (%s)\n", config.App.Name, config.App.Version)
	logger.Info("🌱 Environment: %s\n\n", config.ENV.Env)


	// -------------------------------------------------------------------
	// Embed React UI
	// -------------------------------------------------------------------
	uiFS, err := fs.Sub(embedui.EmbeddedUI, "ui/dist")
	if err != nil {
		log.Fatalf("💀 Web UI Embedding failed: %v", err)
	}
	logger.Info("✅ Web UI ready")
	
	// -------------------------------------------------------------------
	// Initialize Database
	// -------------------------------------------------------------------
	database, err := db.Postgres()
	if err != nil {
		logger.Error("💀 Database connection failed: %v", err)
	}

	
	// -------------------------------------------------------------------
	// Initialize WebSocket Hub
	// -------------------------------------------------------------------
	socketHub := socket.NewHub()
	logger.Info("✅ WebSocket hub ready")


	// -------------------------------------------------------------------
	// Setup Context & Signal Handling (graceful shutdown)
	// -------------------------------------------------------------------
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()



	// -------------------------------------------------------------------
	// Start API Server
	// -------------------------------------------------------------------
	apiServer := server.NewApiServer(config.ENV.Addr, database, uiFS, socketHub)

	logger.Info("🌍 Server listening on http://%s%s\n", GetLocalIP(),config.ENV.Addr)


	// Blocking call — will run until context is cancelled
	if err := apiServer.Run(ctx); err != nil {
		log.Fatalf("Server encountered an error: %v", err)
	}

	logger.Info("Server shutdown complete. Goodbye 👋")
}
