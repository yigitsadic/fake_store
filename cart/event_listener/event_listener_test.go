package event_listener

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/fake_store/cart/database"
	"testing"
)

type mockRepo struct {
	Counter int
}

func (m mockRepo) FindCart(userID string) (*database.Cart, error) {
	panic("implement me")
}

func (m mockRepo) AddToCart(item *database.CartItem) error {
	panic("implement me")
}

func (m mockRepo) RemoveFromCart(itemID, userID string) error {
	panic("implement me")
}

func (m *mockRepo) FlushCart(userID string) {
	m.Counter++
}

func TestEventListener_ListenFlushCartEvents(t *testing.T) {
	repo := &mockRepo{}

	goodMessage := ""
	badMessage := `{"message": "please delete me"}`

	b, err := json.Marshal(flushCartMessage{UserID: "434343"})
	require.Nil(t, err)

	goodMessage = string(b)

	t.Run("it should do nothing with bad message", func(t *testing.T) {
		ch := make(chan *redis.Message)

		listener := &EventListener{
			MessageChan: ch,
			Repository:  repo,
		}

		repo.Counter = 0

		go listener.ListenFlushCartEvents()

		ch <- &redis.Message{
			Channel: ChannelName,
			Payload: badMessage,
		}

		close(ch)

		assert.Equal(t, 0, repo.Counter)
	})

	t.Run("it should call flush cart", func(t *testing.T) {
		ch := make(chan *redis.Message)

		listener := &EventListener{
			MessageChan: ch,
			Repository:  repo,
		}

		repo.Counter = 0

		go listener.ListenFlushCartEvents()

		ch <- &redis.Message{
			Channel: ChannelName,
			Payload: goodMessage,
		}

		close(ch)

		assert.Equal(t, 1, repo.Counter)
	})
}
