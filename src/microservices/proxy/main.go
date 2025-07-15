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
	http.HandleFunc("/api/users", handleMonolith("/api/users"))
	http.HandleFunc("/api/payments", handleMonolith("/api/payments"))
	http.HandleFunc("/api/subscriptions", handleMonolith("/api/subscriptions"))
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
	proxyRequest(w, r, target+"/api/movies")
}

func handleMonolith(apiPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		monolithURL := os.Getenv("MONOLITH_URL")
		fullURL := monolithURL + apiPath
		log.Printf("Proxying %s to: %s", apiPath, fullURL)
		proxyRequest(w, r, fullURL)
	}
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
