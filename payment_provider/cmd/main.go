package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yigitsadic/fake_store/payment_provider/database"
	"github.com/yigitsadic/fake_store/payment_provider/handlers"
	"github.com/yigitsadic/fake_store/payment_provider/trigger"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:5055"
	}

	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/readiness"))

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	repo := database.PaymentIntentRepository{
		Storage: make(map[string]*database.PaymentIntent),
	}

	s := &handlers.Server{
		BaseURL:                 baseURL,
		ShowTemplate:            tmpl,
		PaymentIntentRepository: &repo,
		SendHookRequest:         trigger.SendHookRequest,
	}

	r.Get("/payments/{paymentIntentID}", s.HandleShow())
	r.Post("/payments/complete/{paymentIntentID}", s.HandleComplete())
	r.Post("/payments/create-payment-intent", s.HandleCreate())

	log.Printf("Server is up and running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
