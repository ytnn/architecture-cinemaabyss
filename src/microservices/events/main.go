package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    kafkaBrokers := os.Getenv("KAFKA_BROKERS")
    if kafkaBrokers == "" {
        kafkaBrokers = "localhost:9092"
    }

    go startKafkaConsumer(kafkaBrokers, []string{"movie-events", "user-events", "payment-events"})

    http.HandleFunc("/api/events/health", healthHandler)
    http.HandleFunc("/api/events/movie", createMovieEventHandler)
    http.HandleFunc("/api/events/user", createUserEventHandler)
    http.HandleFunc("/api/events/payment", createPaymentEventHandler)

    log.Println("Events Service listening on :8082")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
