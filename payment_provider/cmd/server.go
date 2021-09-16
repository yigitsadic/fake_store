package main

import (
	"fmt"
	"time"
)

type paymentStatus int

const (
	_ paymentStatus = iota
	paymentInitialized
	paymentCompleted
)

type paymentIntent struct {
	ID          string
	Amount      float64
	ReferenceID string
	Status      paymentStatus

	CreatedAt time.Time

	SuccessURL string
	FailureURL string
	HookURL    string
}

func (i paymentIntent) AmountDisplay() string {
	return fmt.Sprintf("%.2f", i.Amount)
}

var database = make(map[string]paymentIntent)
