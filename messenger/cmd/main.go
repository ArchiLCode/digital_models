package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"messenger/internal/config"
	"messenger/internal/database"
	"messenger/internal/handlers"
	"messenger/internal/migrations"
	"messenger/internal/repository"
	"messenger/internal/websocket"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize database connection
	if err := database.InitDB(cfg); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Run database migrations
	migrationsDir := filepath.Join("migrations")
	if err := migrations.RunMigrations(database.GetDB(), migrationsDir); err != nil {
		log.Fatal("Error running migrations:", err)
	}
	log.Println("Database migrations completed successfully")
	// Initialize repository
	repo := repository.NewRepository(database.GetDB())

	// Create WebSocket hub with repository
	hub := websocket.NewHub(repo)
	go hub.Run()

	// Initialize handlers
	handler := handlers.NewHandler(repo, hub)

	// Create router and register routes
	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Server starting on port %d", cfg.ServerPort)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
