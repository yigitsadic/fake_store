package event_handlers

import (
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/products/database"
)

const (
	ProductInfoPopulateChannel = "PRODUCT_INFO_POPULATE_CHANNEL"
	ProductPopulatedChannel    = "PRODUCT_POPULATED_CHANNEL"
)

type EventHandler struct {
	PopulateCartItemFunc func(product database.Product)
	MessageChan          <-chan *redis.Message
	ProductRepository    database.Repository
}

func (h *EventHandler) ListenProductPopulateMessages() {
	for msg := range h.MessageChan {
		message, err := unmarshalProductDataRequest(msg.Payload)
		if err != nil {
			continue
		}

		if message.ProductID != "" && message.CartItemID != "" {
			product, err := h.ProductRepository.FetchOne(message.ProductID)
			if err != nil {
				continue
			}

			h.PopulateCartItemFunc(*product)
		}
	}
}
