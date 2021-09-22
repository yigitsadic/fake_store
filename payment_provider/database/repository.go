package database

import (
	"errors"
	"fmt"
	"time"
)

type paymentStatus int

const (
	_ paymentStatus = iota

	// PaymentInitialized means process started. It's default state.
	PaymentInitialized

	// PaymentCompleted means process completed.
	PaymentCompleted
)

var errorRecordNotFound = errors.New("record not found on database")

// PaymentHookMessage contains information that will send to hook.
type PaymentHookMessage struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// PaymentIntent represents payments both completed and started.
type PaymentIntent struct {
	ID          string        `bson:"_id,omitempty"`
	Amount      float64       `bson:"amount,omitempty"`
	ReferenceID string        `bson:"reference_id,omitempty"`
	Status      paymentStatus `bson:"status,omitempty"`

	CreatedAt time.Time `bson:"created_at,omitempty"`

	SuccessURL string `bson:"success_url,omitempty"`
	FailureURL string `bson:"failure_url,omitempty"`
	HookURL    string `bson:"hook_url,omitempty"`
}

// CreateHookMessage creates PaymentHookMessage from intent.
func (i PaymentIntent) CreateHookMessage() PaymentHookMessage {
	return PaymentHookMessage{
		ID:          i.ID,
		Amount:      i.Amount,
		ReferenceID: i.ReferenceID,
		Status:      int(i.Status),
		CreatedAt:   i.CreatedAt,
	}
}

// AmountDisplay displays float as two decimal string.
func (i PaymentIntent) AmountDisplay() string {
	return fmt.Sprintf("%.2f", i.Amount)
}

// Repository is an interface for interacting between application and database.
type Repository interface {
	Create(referenceID, hookURL, successURL, failureURL string, amount float64) (*PaymentIntent, error)
	FindOne(ID string) (*PaymentIntent, error)
	MarkAsCompleted(ID string) error
}
