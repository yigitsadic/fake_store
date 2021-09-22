package database

import (
	"github.com/bxcodec/faker/v3"
	"time"
)

// MockPaymentIntentRepository is in-memory, database mimicking struct.
type MockPaymentIntentRepository struct {
	Storage map[string]*PaymentIntent
}

// Create inserts new intent to in-memory database with given parameters.
func (p *MockPaymentIntentRepository) Create(referenceID, hookURL, successURL, failureURL string, amount float64) (*PaymentIntent, error) {
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

// FindOne fetches record from in-memory database with given ID.
func (p MockPaymentIntentRepository) FindOne(ID string) (*PaymentIntent, error) {
	record, ok := p.Storage[ID]
	if !ok {
		return nil, errorRecordNotFound
	}

	return record, nil
}

// MarkAsCompleted marks payment as complete for given ID.
func (p *MockPaymentIntentRepository) MarkAsCompleted(ID string) error {
	record, ok := p.Storage[ID]
	if !ok {
		return errorRecordNotFound
	}

	record.Status = PaymentCompleted

	return nil
}
