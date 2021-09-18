package event_listener

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/database"
)

const ChannelName = "FLUSH_CART_CHANNEL"

type flushCartMessage struct {
	UserID string `json:"user_id"`
}

type EventListener struct {
	RedisClient *redis.Client
	Ctx         context.Context
	Repository  database.Repository
}

func (l *EventListener) ListenFlushCartEvents() {
	pubSub := l.RedisClient.Subscribe(l.Ctx, ChannelName)

	ch := pubSub.Channel()

	for msg := range ch {
		var cartMessage flushCartMessage

		err := json.Unmarshal([]byte(msg.Payload), &cartMessage)
		if err == nil {
			l.Repository.FlushCart(cartMessage.UserID)
		}
	}
}
