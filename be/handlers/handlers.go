package handlers

import (
	"andersonlira.com/app-api/db"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Response represents the standard API response format
type Response struct {
	Data      interface{} `json:"data,omitempty"`
	Error     *string     `json:"error,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// HealthCheck handles the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check database connectivity
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "ok"
	if err := db.Pool.Ping(ctx); err != nil {
		dbStatus = "unavailable"
	}

	data := map[string]interface{}{
		"status":   "ok",
		"service":  "app-api",
		"version":  "1.0.0",
		"database": dbStatus,
	}

	writeJSON(w, data, http.StatusOK)
}

// AppConfig represents application configuration for API responses
type AppConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// GetConfigHandler handles the config retrieval endpoint
func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	dbConfig, err := db.GetAppConfig(ctx)
	if err != nil {
		writeError(w, "Failed to fetch configuration", http.StatusInternalServerError)
		return
	}

	// Convert db.AppConfig to handlers.AppConfig
	config := AppConfig{
		Name:    dbConfig.Name,
		Version: dbConfig.Version,
	}

	writeJSON(w, config, http.StatusOK)
}

// Helper functions

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	response := Response{
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, message string, status int) {
	response := Response{
		Error:     &message,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
