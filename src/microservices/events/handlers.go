package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]bool{"status": true})
}

func createMovieEventHandler(w http.ResponseWriter, r *http.Request) {
    var payload MovieEvent
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid movie event", http.StatusBadRequest)
        return
    }

    event := Event{
        ID:        fmt.Sprintf("movie-%d-%s", payload.MovieID, payload.Action),
        Type:      "movie",
        Timestamp: time.Now().UTC().Format(time.RFC3339),
        Payload:   payload,
    }

    publishAndRespond(w, "movie-events", event)
}

func createUserEventHandler(w http.ResponseWriter, r *http.Request) {
    var payload UserEvent
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid user event", http.StatusBadRequest)
        return
    }

    event := Event{
        ID:        fmt.Sprintf("user-%d-%s", payload.UserID, payload.Action),
        Type:      "user",
        Timestamp: payload.Timestamp,
        Payload:   payload,
    }

    publishAndRespond(w, "user-events", event)
}

func createPaymentEventHandler(w http.ResponseWriter, r *http.Request) {
    var payload PaymentEvent
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid payment event", http.StatusBadRequest)
        return
    }

    event := Event{
        ID:        fmt.Sprintf("payment-%d-%s", payload.PaymentID, payload.Status),
        Type:      "payment",
        Timestamp: payload.Timestamp,
        Payload:   payload,
    }

    publishAndRespond(w, "payment-events", event)
}
