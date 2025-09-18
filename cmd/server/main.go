package main

import (
	"context"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"
	embedui "watchtower"
	"watchtower/config"
	"watchtower/db"
	"watchtower/internal/server"
)

func main() {
	log.Printf("Starting Server...\n\n",)

	// Prepare frontend FS
	uiFS, e := fs.Sub(embedui.EmbeddedUI, "ui/dist")
	if e != nil {
		log.Fatal(e)
	}

	// DB connection
	var db, err = db.SQLite()
	if err != nil {log.Fatal(err)}


	// Handle interrupts
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := server.NewApiServer(config.ENV.Addr, db, uiFS)
	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
