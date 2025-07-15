package main

import (
    "context"
    "encoding/json"
    "net/http" 
    "log"
    "os"
    "time"

    "github.com/segmentio/kafka-go"
)

func publishAndRespond(w http.ResponseWriter, topic string, event Event) {
    data, err := json.Marshal(event)
    if err != nil {
        http.Error(w, "Failed to encode event", http.StatusInternalServerError)
        return
    }

    writer := kafka.Writer{
        Addr:     kafka.TCP(os.Getenv("KAFKA_BROKERS")),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},
    }
    defer writer.Close()

    msg := kafka.Message{Value: data}
    err = writer.WriteMessages(context.Background(), msg)
    if err != nil {
        http.Error(w, "Failed to publish", http.StatusInternalServerError)
        return
    }

    log.Printf("Published to %s: %s\n", topic, data)

    resp := EventResponse{
        Status:    "success",
        Partition: 0,
        Offset:    0,
        Event:     event,
    }

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func startKafkaConsumer(brokers string, topics []string) {
    for _, topic := range topics {
        go func(t string) {
            reader := kafka.NewReader(kafka.ReaderConfig{
                Brokers:  []string{brokers},
                GroupID:  "events-consumer-" + t,
                Topic:    t,
                MinBytes: 1,
                MaxBytes: 10e6,
            })
            defer reader.Close()

            for {
                m, err := reader.ReadMessage(context.Background())
                if err != nil {
                    log.Printf("Error reading from topic %s: %v", t, err)
                    time.Sleep(1 * time.Second)
                    continue
                }
                log.Printf("Consumed from %s: %s\n", t, string(m.Value))
            }
        }(topic)
    }
}
