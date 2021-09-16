package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
	"time"
)

type paymentStatus int

const (
	_ paymentStatus = iota
	paymentPending
	paymentCompleted
)

type paymentIntentMessage struct {
	ID          string        `json:"id"`
	Amount      float64       `json:"amount"`
	ReferenceID string        `json:"reference_id"`
	Status      paymentStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	PaymentURL  string        `json:"payment_url"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4035"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("Unable to connect redis")
	}

	r := chi.NewRouter()

	r.Post("/api/payment/webhooks", hookHandler(rdb))

	log.Printf("Server is up and running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
