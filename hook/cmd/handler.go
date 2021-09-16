package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"io"
	"net/http"
)

const channelName = "PAYMENTS_COMPLETE_CHANNEL"

type redisClient interface {
	Publish(context.Context, string, interface{}) *redis.IntCmd
}

func hookHandler(rds redisClient) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var message paymentIntentMessage
		b, err := io.ReadAll(request.Body)
		defer request.Body.Close()

		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = json.Unmarshal(b, &message)
		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		if message.Status != paymentCompleted {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = rds.Publish(request.Context(), channelName, string(b)).Err()
		if err != nil {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
