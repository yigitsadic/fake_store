package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func buildRequest(intent paymentIntent) (*http.Request, error) {
	message := &paymentIntentMessage{
		ID:          intent.ID,
		Amount:      intent.Amount,
		ReferenceID: intent.ReferenceID,
		Status:      intent.Status,
		CreatedAt:   intent.CreatedAt,
	}

	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, intent.HookURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func sendPaymentRequestToHookURL(intent paymentIntent) error {
	req, err := buildRequest(intent)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: time.Second * 30}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New("unable to send payment successful hook")
}
