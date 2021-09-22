package events

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/database"
)

const (
	// FlushCartChannelName stores redis channel name for flush cart items.
	FlushCartChannelName = "FLUSH_CART_CHANNEL"

	// CartItemPopulatedChannelName is channel name for data populated from products service.
	CartItemPopulatedChannelName = "CART_ITEM_POPULATED"
)

type flushCartMessage struct {
	UserID string `json:"user_id"`
}

// EventListener is a struct which contains flush cart and product data populated messages channels.
type EventListener struct {
	FlushCartMessageChan     <-chan *redis.Message
	ProductInfoPopulatedChan <-chan *redis.Message
	Repository               database.Repository
}

// ListenFlushCartEvents listens messages from given channel and if payload matches
// flushes cart content.
func (l *EventListener) ListenFlushCartEvents() {
	for msg := range l.FlushCartMessageChan {
		var cartMessage flushCartMessage

		err := json.Unmarshal([]byte(msg.Payload), &cartMessage)

		if err == nil && cartMessage.UserID != "" {
			l.Repository.FlushCart(cartMessage.UserID)
		}
	}
}

// ListenCartItemDataPopulationMessages listens messages from product data populated channel.
// Updates cart item with populated product data.
func (l *EventListener) ListenCartItemDataPopulationMessages() {
	for msg := range l.ProductInfoPopulatedChan {
		var prodMessage database.CartItemProductMessage

		err := json.Unmarshal([]byte(msg.Payload), &prodMessage)
		if err == nil && prodMessage.ProductID != "" && prodMessage.CartItemID != "" {
			l.Repository.UpdateCartItem(prodMessage)
		}
	}
}
