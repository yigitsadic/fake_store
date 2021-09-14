package graph

import (
	"net/http"
	"time"
)

type paymentStatus int

const (
	_ paymentStatus = iota
	paymentInitialized
	paymentCompleted
)

type createPaymentIntentRequestRequest struct {
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

func createPaymentIntentRequest() (*http.Request, error) {
	return nil, nil
}

func createPaymentIntent(userID string, paymentAmount float64) error {
	return nil
}
