package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var baseURL string

func init() {
	database["yigit"] = paymentIntent{
		ID:          "yigit",
		Amount:      15.25,
		ReferenceID: "1231245245",
		Status:      paymentInitialized,
		CreatedAt:   time.Now().UTC(),
		SuccessURL:  "http://localhost:3000/orders?success",
		FailureURL:  "http://localhost:3000/orders?success",
		HookURL:     "http://localhost:3000/orders?success",
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}

	baseURL = os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:5055"
	}

	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)

	r.Use(middleware.Heartbeat("/readiness"))

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	r.Get("/payments/{paymentIntentID}", handleShowPaymentIntent(tmpl))
	r.Post("/payments/complete/{paymentIntentID}", handleCompletePaymentIntent)

	r.Post("/payments/create-payment-intent", handleCreatePaymentIntent)

	log.Printf("Server is up and running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
