package event_handlers

import (
	"encoding/json"
	"errors"
)

const (
	PAYMENT_COMPLETE_CHANNEL = "PAYMENTS_COMPLETE_CHANNEL"
	FLUSH_CART_CHANNEL       = "FLUSH_CART_CHANNEL"
)

type paymentIntentMessage struct {
	ReferenceID string `json:"reference_id"`
}

type flushCartMessage struct {
	UserID string `json:"user_id"`
}

// unmarshalMessage parses reference id from given string.
func unmarshalMessage(given string) (string, error) {
	var paymentMessage paymentIntentMessage

	err := json.Unmarshal([]byte(given), &paymentMessage)
	if err != nil {
		return "", err
	}

	if paymentMessage.ReferenceID == "" {
		return "", errors.New("reference ID not found on payload")
	}

	return paymentMessage.ReferenceID, nil
}
