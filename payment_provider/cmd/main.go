package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"html/template"
	"log"
	"net/http"
	"os"
)

var baseURL string

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
