package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type paymentStatus int

const (
	_ paymentStatus = iota
	_
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

	r.Post("/api/payment/webhooks", func(writer http.ResponseWriter, request *http.Request) {
		var message paymentIntentMessage
		b, err := io.ReadAll(request.Body)
		defer request.Body.Close()

		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = json.Unmarshal(b, &message)
		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if message.Status != paymentCompleted {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = rdb.Publish(request.Context(), "PAYMENTS_COMPLETE_CHANNEL", string(b)).Err()
		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		writer.WriteHeader(http.StatusOK)
	})

	log.Printf("Server is up and running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
