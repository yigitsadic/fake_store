package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type paymentIntentRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceID string  `json:"reference_id"`

	HookURL    string `json:"hook_url"`
	SuccessURL string `json:"success_url"`
	FailureURL string `json:"failure_url"`
}

type paymentIntentResponse struct {
	ID         string `json:"id"`
	PaymentURL string `json:"payment_url"`
}

func buildRequest(ctx context.Context, target string, message paymentIntentRequest) (*http.Request, error) {
	b, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func createPaymentIntent(ctx context.Context, target string, message paymentIntentRequest) (string, error) {
	req, err := buildRequest(ctx, target, message)
	if err != nil {
		return "", err
	}

	client := &http.Client{Timeout: time.Second * 30}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	var resp paymentIntentResponse

	err = json.Unmarshal(content, &resp)
	if err != nil {
		return "", err
	}

	if resp.PaymentURL != "" {
		return resp.PaymentURL, nil
	}

	return "", errors.New("unable to send payment intent create request to payment provider")
}
