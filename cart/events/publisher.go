package events

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/database"
)

const (
	// PopulateCartItemChannelName channel name for cart item product info population requests.
	PopulateCartItemChannelName = "POPULATE_CART_ITEM"
)

type redisKind interface {
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
}

// PublishToRedis publishes product data need to populate message to Redis pub/sub.
func PublishToRedis(rds redisKind, cartItemID, productID string) {
	b, err := json.Marshal(database.CartItemProductMessage{ProductID: productID, CartItemID: cartItemID})
	if err == nil {
		rds.Publish(context.Background(), PopulateCartItemChannelName, string(b))
	}
}
