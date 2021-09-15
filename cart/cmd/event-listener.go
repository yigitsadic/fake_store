package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type flushCartMessage struct {
	UserID string `json:"user_id"`
}

type eventListener struct {
	RedisClient *redis.Client
	Ctx         context.Context
	Database    *cartDatabase
}

func (l *eventListener) ListenFlushCartEvents() {
	pubSub := l.RedisClient.Subscribe(l.Ctx, "FLUSH_CART_CHANNEL")

	ch := pubSub.Channel()

	for msg := range ch {
		var cartMessage flushCartMessage

		err := json.Unmarshal([]byte(msg.Payload), &cartMessage)
		if err == nil {
			l.Database.Storage[cartMessage.UserID] = []cartItem{}
		}
	}
}
