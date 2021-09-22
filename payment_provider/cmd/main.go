package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yigitsadic/fake_store/payment_provider/database"
	"github.com/yigitsadic/fake_store/payment_provider/handlers"
	"github.com/yigitsadic/fake_store/payment_provider/trigger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
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

	// Mongo connection.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://database:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err)
	}

	// Inject Mongo client to repository.
	repo := database.PaymentIntentRepository{
		Storage: client.Database("fake_store"),
		Ctx:     context.Background(),
	}

	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/readiness"))

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

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
