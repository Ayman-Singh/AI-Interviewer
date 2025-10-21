package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ai-interviewer/backend/internal/ai"
	"github.com/ai-interviewer/backend/internal/config"
	"github.com/ai-interviewer/backend/internal/database"
	"github.com/ai-interviewer/backend/internal/handlers"
	"github.com/ai-interviewer/backend/internal/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("DEBUG: Allowed Origins Loaded: %v", cfg.AllowedOrigins)
	// Connect to database with retry
	var db *database.Database
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = database.New(cfg)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
	}
	defer db.Close()

	log.Println("Successfully connected to database")

	// Initialize AI service
	aiService, err := ai.NewAIService(cfg.GeminiAPIKey)
	if err != nil {
		log.Fatalf("Failed to initialize AI service: %v", err)
	}

	// Initialize repository and handlers
	repo := repository.New(db.DB)
	handler := handlers.New(repo, aiService)

	// Setup router
	router := mux.NewRouter()

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(c.Handler)

	// Logging middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// API routes
	api := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	api.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	api.HandleFunc("/interview/start", handler.StartInterview).Methods("POST")
	api.HandleFunc("/interview/{id}", handler.GetInterview).Methods("GET")
	api.HandleFunc("/interview/submit", handler.SubmitAnswer).Methods("POST")
	api.HandleFunc("/interviews", handler.GetUserInterviews).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}