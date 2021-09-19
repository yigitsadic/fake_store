package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type successfulMockRedisClient struct {
	CallTime int
	Channel  string
}

func (m *successfulMockRedisClient) Publish(ctx context.Context, s string, i interface{}) *redis.IntCmd {
	m.CallTime++
	m.Channel = s

	return &redis.IntCmd{}
}

type failureMockRedisClient struct {
	CallTime int
}

func (m *failureMockRedisClient) Publish(context.Context, string, interface{}) *redis.IntCmd {
	m.CallTime++

	res := &redis.IntCmd{}
	res.SetErr(errors.New("unable to publish message"))

	return res
}

func Test_handleBookHandler(t *testing.T) {
	t.Run("it will return error for unsupported payload", func(t *testing.T) {
		client := http.Client{}
		r := chi.NewRouter()
		c := createSuccessfulMockRedisClient()

		r.Post("/", hookHandler(c))
		ts := httptest.NewServer(r)
		defer ts.Close()

		req := buildBadRequest(ts.URL)
		res, err := client.Do(req)

		assert.Nil(t, err, "unexpected to get an error")
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "expected to get 422 response")
	})

	t.Run("it will return error for incomplete payment status", func(t *testing.T) {
		client := http.Client{}
		r := chi.NewRouter()
		c := createSuccessfulMockRedisClient()

		r.Post("/", hookHandler(c))
		ts := httptest.NewServer(r)
		defer ts.Close()

		req := buildPendingPaymentStatusRequest(ts.URL)
		res, err := client.Do(req)

		assert.Nil(t, err, "unexpected to get an error")
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "expected to get 422 response")
	})

	t.Run("it will return error if it cannot publish message to redis", func(t *testing.T) {
		client := http.Client{}
		r := chi.NewRouter()
		c := createFailureMockRedisClient()

		r.Post("/", hookHandler(c))
		ts := httptest.NewServer(r)
		defer ts.Close()

		req := buildGoodRequest(ts.URL)
		res, err := client.Do(req)

		assert.Nil(t, err, "unexpected to get an error")
		assert.Equal(t, 1, c.CallTime, "expected to called for once")
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode, "expected to get 422 response")
	})

	t.Run("it will return status ok if everything went well", func(t *testing.T) {
		client := http.Client{}
		r := chi.NewRouter()
		c := createSuccessfulMockRedisClient()

		r.Post("/", hookHandler(c))
		ts := httptest.NewServer(r)
		defer ts.Close()

		req := buildGoodRequest(ts.URL)
		res, err := client.Do(req)

		assert.Nil(t, err, "unexpected to get an error")
		assert.Equal(t, 1, c.CallTime, "expected to get called for once")
		assert.Equalf(t, channelName, c.Channel, "expected channel name was %s but got %s", channelName, c.Channel)
		assert.Equal(t, http.StatusOK, res.StatusCode, "expected to get 200 response")
	})
}

func buildGoodRequest(target string) *http.Request {
	b, _ := json.Marshal(paymentIntentMessage{
		ID:          "qweqweqweq",
		Amount:      65.12,
		ReferenceID: "1312313132",
		Status:      paymentCompleted,
		CreatedAt:   time.Now(),
		PaymentURL:  "",
	})

	req, _ := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req
}

func buildBadRequest(target string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, target, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req
}

func buildPendingPaymentStatusRequest(target string) *http.Request {
	b, _ := json.Marshal(paymentIntentMessage{
		ID:          "qweqweqweq",
		Amount:      65.12,
		ReferenceID: "1312313132",
		Status:      paymentPending,
		CreatedAt:   time.Now(),
		PaymentURL:  "",
	})

	req, _ := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req
}

func createSuccessfulMockRedisClient() *successfulMockRedisClient {
	c := successfulMockRedisClient{}

	return &c
}

func createFailureMockRedisClient() *failureMockRedisClient {
	c := failureMockRedisClient{}

	return &c
}
