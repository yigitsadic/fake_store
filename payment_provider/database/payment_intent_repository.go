package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// PaymentIntentRepository Mongo integrated struct to handle create, find and update as defined in interface.
type PaymentIntentRepository struct {
	Storage *mongo.Database
	Ctx     context.Context
}

// Create creates document in mongo collection.
func (p *PaymentIntentRepository) Create(referenceID, hookURL, successURL, failureURL string, amount float64) (*PaymentIntent, error) {
	intent := PaymentIntent{
		Amount:      amount,
		ReferenceID: referenceID,
		Status:      PaymentInitialized,
		CreatedAt:   time.Now().UTC(),
		SuccessURL:  successURL,
		FailureURL:  failureURL,
		HookURL:     hookURL,
	}

	result, err := p.Storage.Collection("payment_intents").InsertOne(p.Ctx, intent)
	if err != nil {
		return nil, err
	}

	intent.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return &intent, nil
}

// FindOne retrieves one record.
func (p *PaymentIntentRepository) FindOne(ID string) (*PaymentIntent, error) {
	var intent PaymentIntent

	err := p.Storage.Collection("payment_intents").FindOne(p.Ctx, bson.M{"_id": ID}).Decode(&intent)
	if err != nil {
		return nil, err
	}

	return &intent, err
}

// MarkAsCompleted updates given id record as completed.
func (p *PaymentIntentRepository) MarkAsCompleted(ID string) error {
	result, err := p.Storage.Collection("payment_intents").UpdateOne(
		p.Ctx,
		bson.M{"_id": ID},
		bson.M{"status": PaymentCompleted},
	)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("could not update record")
	}

	return nil
}
