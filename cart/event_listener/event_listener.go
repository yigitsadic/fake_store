package event_listener

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/database"
)

const ChannelName = "FLUSH_CART_CHANNEL"

type flushCartMessage struct {
	UserID string `json:"user_id"`
}

type EventListener struct {
	MessageChan <-chan *redis.Message
	Repository  database.Repository
}

func (l *EventListener) ListenFlushCartEvents() {
	for msg := range l.MessageChan {
		var cartMessage flushCartMessage

		err := json.Unmarshal([]byte(msg.Payload), &cartMessage)

		if err == nil && cartMessage.UserID != "" {
			l.Repository.FlushCart(cartMessage.UserID)
		}
	}
}
