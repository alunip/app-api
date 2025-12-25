package main

import (
	"andersonlira.com/app-api/db"
	"andersonlira.com/app-api/handlers"
	"andersonlira.com/app-api/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create context for initialization
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Load database configuration
	dbConfig := db.LoadConfig()

	// Run migrations
	log.Println("Running database migrations...")
	if err := db.RunMigrations(dbConfig); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Connect to database
	log.Println("Connecting to database...")
	pool, err := db.Connect(ctx, dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.Pool = pool
	defer db.Close()

	// Get port from environment or default to 8080
	port := getEnv("PORT", "8080")

	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/api/health", handlers.HealthCheck)
	mux.HandleFunc("/api/config", handlers.GetConfigHandler)

	// Wrap with CORS middleware
	handler := middleware.CORS(mux)

	// Create server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create context with timeout for shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
