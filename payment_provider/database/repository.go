package database

import (
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"time"
)

type paymentStatus int

const (
	_ paymentStatus = iota
	PaymentInitialized
	PaymentCompleted
)

var errorRecordNotFound = errors.New("record not found on database")

type PaymentHookMessage struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaymentIntent struct {
	ID          string
	Amount      float64
	ReferenceID string
	Status      paymentStatus

	CreatedAt time.Time

	SuccessURL string
	FailureURL string
	HookURL    string
}

func (i PaymentIntent) CreateHookMessage() PaymentHookMessage {
	return PaymentHookMessage{
		ID:          i.ID,
		Amount:      i.Amount,
		ReferenceID: i.ReferenceID,
		Status:      int(i.Status),
		CreatedAt:   i.CreatedAt,
	}
}

func (i PaymentIntent) AmountDisplay() string {
	return fmt.Sprintf("%.2f", i.Amount)
}

type Repository interface {
	Create(referenceID, hookURL, successURL, failureURL string, amount float64) (*PaymentIntent, error)
	FindOne(ID string) (*PaymentIntent, error)
	MarkAsCompleted(ID string) error
}

type PaymentIntentRepository struct {
	Storage map[string]*PaymentIntent
}

func (p *PaymentIntentRepository) Create(referenceID, hookURL, successURL, failureURL string, amount float64) (*PaymentIntent, error) {
	record := PaymentIntent{
		ID:          faker.UUIDHyphenated(),
		Amount:      amount,
		ReferenceID: referenceID,
		Status:      PaymentInitialized,
		CreatedAt:   time.Now().UTC(),
		SuccessURL:  successURL,
		FailureURL:  failureURL,
		HookURL:     hookURL,
	}

	return &record, nil
}

func (p PaymentIntentRepository) FindOne(ID string) (*PaymentIntent, error) {
	record, ok := p.Storage[ID]
	if !ok {
		return nil, errorRecordNotFound
	}

	return record, nil
}

func (p *PaymentIntentRepository) MarkAsCompleted(ID string) error {
	record, ok := p.Storage[ID]
	if !ok {
		return errorRecordNotFound
	}

	record.Status = PaymentCompleted

	return nil
}
