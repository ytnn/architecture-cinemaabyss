package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/api/movies", handleMovies)
	http.HandleFunc("/api/users", handleMonolith)
	http.HandleFunc("/api/payments", handleMonolith)
	http.HandleFunc("/api/subscriptions", handleMonolith)
	http.HandleFunc("/health", handleHealth)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Proxy server running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleMovies(w http.ResponseWriter, r *http.Request) {
	target := selectMoviesTarget()
	log.Printf("Proxying /api/movies to: %s", target)
	proxyRequest(w, r, target)
}

func handleMonolith(w http.ResponseWriter, r *http.Request) {
	monolithURL := os.Getenv("MONOLITH_URL")
	log.Printf("Proxying %s to %s", r.URL.RequestURI(), monolithURL)
	proxyRequest(w, r, monolithURL)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}

func selectMoviesTarget() string {
	monolithURL := os.Getenv("MONOLITH_URL")
	moviesURL := os.Getenv("MOVIES_SERVICE_URL")
	gradualMigration := os.Getenv("GRADUAL_MIGRATION") == "true"

	if !gradualMigration {
		return monolithURL
	}

	percent := getMigrationPercent()
	rand.Seed(time.Now().UnixNano())

	if rand.Intn(100) < percent {
		return moviesURL
	}
	return monolithURL
}
