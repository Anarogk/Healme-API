package main

import (
	"log"

	"github.com/Anarogk/healthcare-api/configs"
	"github.com/Anarogk/healthcare-api/internal/api"
	"github.com/Anarogk/healthcare-api/internal/db"
)

func main() {
	// Load configurations
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := db.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize and start the server
	server := api.NewServer(config, db)
	err = server.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
