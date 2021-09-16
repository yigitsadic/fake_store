package main

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/cenkalti/backoff/v4"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"time"
)

type createPaymentIntentRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceID string  `json:"reference_id"`

	HookURL    string `json:"hook_url"`
	SuccessURL string `json:"success_url"`
	FailureURL string `json:"failure_url"`
}

type paymentIntentMessage struct {
	ID          string        `json:"id"`
	Amount      float64       `json:"amount"`
	ReferenceID string        `json:"reference_id"`
	Status      paymentStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	PaymentURL  string        `json:"payment_url"`
}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	b := createPaymentIntentRequest{}

	json.NewDecoder(request.Body).Decode(&b)

	intent := paymentIntent{
		ID:          faker.UUIDHyphenated(),
		Amount:      b.Amount,
		ReferenceID: b.ReferenceID,
		Status:      paymentInitialized,
		CreatedAt:   time.Now().UTC(),
		HookURL:     b.HookURL,
		SuccessURL:  b.SuccessURL,
		FailureURL:  b.FailureURL,
	}

	database[intent.ID] = intent

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(&paymentIntentMessage{
		ID:          intent.ID,
		Amount:      intent.Amount,
		ReferenceID: intent.ReferenceID,
		Status:      intent.Status,
		CreatedAt:   intent.CreatedAt,
		PaymentURL:  baseURL + "/payments/" + intent.ID,
	})
}

func handleShowPaymentIntent(tmp *template.Template) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		paymentIntentID := chi.URLParam(request, "paymentIntentID")

		record, ok := database[paymentIntentID]

		if ok {
			tmp.Execute(writer, &record)
		} else {
			writer.WriteHeader(http.StatusNotFound)

			return
		}
	}
}

func handleCompletePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	paymentIntentID := chi.URLParam(request, "paymentIntentID")

	record, ok := database[paymentIntentID]
	targetURL := record.FailureURL

	if ok {
		newRecord := paymentIntent{
			ID:          record.ID,
			Amount:      record.Amount,
			ReferenceID: record.ReferenceID,
			Status:      paymentCompleted,
			CreatedAt:   record.CreatedAt,
			HookURL:     record.HookURL,
			SuccessURL:  record.SuccessURL,
			FailureURL:  record.FailureURL,
		}

		operation := func() error {
			return sendPaymentRequestToHookURL(newRecord)
		}

		// send successful payment hook
		err := backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
		if err != nil {
			http.Redirect(writer, request, targetURL, http.StatusFound)
			return
		}

		database[newRecord.ID] = newRecord

		targetURL = record.SuccessURL
	}

	http.Redirect(writer, request, targetURL, http.StatusFound)
}
