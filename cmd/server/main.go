package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"watchtower/config"
	"watchtower/db"
	"watchtower/internal/server"
)


func main() {
	log.Printf("Starting Server...\n\n",)

	// DB connection
	var db, err = db.SQLite()
	if err != nil {log.Fatal(err)}


	// Handle interrupts
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := server.NewApiServer(config.ENV.Port, db)
	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
