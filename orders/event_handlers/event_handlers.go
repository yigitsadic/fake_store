package event_handlers

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/orders/database"
)

type EventHandler struct {
	FlushCartFunc   func(message string)
	MessageChan     <-chan *redis.Message
	OrderRepository database.Repository
}

func (h *EventHandler) ListenPaymentCompleteEvents() {
	for msg := range h.MessageChan {
		orderID, err := unmarshalMessage(msg.Payload)
		if err == nil {
			userID, err := h.OrderRepository.Complete(orderID)
			if err == nil {
				h.FlushCartFunc(fmt.Sprintf("{%q: %q}", "user_id", userID))
			}
		}
	}
}
