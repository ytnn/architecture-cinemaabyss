package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/movies", handleMovies)
	mux.HandleFunc("/api/users", handleUsers)
	mux.HandleFunc("/health", handleHealth)

	log.Println("🚀 Proxy started at port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
