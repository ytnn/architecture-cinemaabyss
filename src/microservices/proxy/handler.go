package main

import (
	"math/rand"
	"net/http"
	"log"
	"os"  
	"encoding/json"
	"time"
)

func handleMovies(w http.ResponseWriter, r *http.Request) {
	monolithURL := os.Getenv("MONOLITH_URL")
	moviesURL := os.Getenv("MOVIES_SERVICE_URL")
	gradualMigration := os.Getenv("GRADUAL_MIGRATION") == "true"

	migrationPercent := getMigrationPercent()

	rand.Seed(time.Now().UnixNano())
	useNew := rand.Intn(100) < migrationPercent

	var target string
	if gradualMigration && useNew {
		target = moviesURL + "/api/movies"
		log.Println("Service")
	} else {
		target = monolithURL + "/api/movies"
		log.Println("Monolith")
	}

	proxyRequest(w, r, target)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	monolithURL := os.Getenv("MONOLITH_URL")
	proxyRequest(w, r, monolithURL+"/api/users")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}

