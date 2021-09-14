package main

import "time"

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

	HookURL         string
	HookCredentials string
}

var dataBase = make(map[string]paymentIntent)
