package events

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/cart/database"
	"testing"
)

type mockRedis struct {
	CallCounter int
	Chan        string
	Payload     string
}

func (m *mockRedis) Publish(_ context.Context, channel string, message interface{}) *redis.IntCmd {
	m.Payload = message.(string)
	m.Chan = channel
	m.CallCounter++

	return &redis.IntCmd{}
}

func (m *mockRedis) FormattedPayload() database.CartItemProductMessage {
	message := database.CartItemProductMessage{}
	json.Unmarshal([]byte(m.Payload), &message)

	return message
}

func TestPublishToRedis(t *testing.T) {
	rds := new(mockRedis)

	PublishToRedis(rds, "lorem", "ipsum")

	for rds.CallCounter == 0 {
	}

	assert.Equal(t, PopulateCartItemChannelName, rds.Chan)
	assert.Equal(t, 1, rds.CallCounter)
	assert.Equal(t, "lorem", rds.FormattedPayload().CartItemID)
	assert.Equal(t, "ipsum", rds.FormattedPayload().ProductID)
}
