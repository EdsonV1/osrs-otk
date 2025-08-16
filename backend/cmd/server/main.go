package main

import (
	"log"

	"osrs-xp-kits/internal/config"
	"osrs-xp-kits/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create and start server
	srv := server.New(cfg)

	log.Printf("Starting server on %s", cfg.GetAddr())
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
