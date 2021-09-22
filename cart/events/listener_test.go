package events

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/cart/database"
	"testing"
	"time"
)

var counter int
var gotUserID string

type mockRepo struct {
}

func (m mockRepo) UpdateCartItem(message database.CartItemProductMessage) {
	panic("implement me")
}

func (m mockRepo) FindCart(userID string) (*database.Cart, error) {
	panic("implement me")
}

func (m mockRepo) AddToCart(userID string, productID string) (string, error) {
	panic("implement me")
}

func (m mockRepo) RemoveFromCart(itemID, userID string) error {
	panic("implement me")
}

func (m *mockRepo) FlushCart(userID string) {
	counter++
	gotUserID = userID
}

func TestEventListener_ListenFlushCartEvents(t *testing.T) {
	repo := &mockRepo{}

	goodMessage := `{"user_id": "434343"}`
	badMessage := `{"message": "please delete me"}`

	t.Run("it should do nothing with bad message", func(t *testing.T) {
		gotUserID = ""
		counter = 0

		ch := make(chan *redis.Message)

		listener := &EventListener{
			FlushCartMessageChan: ch,
			Repository:           repo,
		}

		go listener.ListenFlushCartEvents()

		ch <- &redis.Message{
			Channel: FlushCartChannelName,
			Payload: badMessage,
		}

		close(ch)

		assert.Equal(t, 0, counter)
		assert.Equal(t, "", gotUserID)
	})

	t.Run("it should call flush cart", func(t *testing.T) {
		gotUserID = ""
		counter = 0

		ch := make(chan *redis.Message)

		listener := &EventListener{
			FlushCartMessageChan: ch,
			Repository:           repo,
		}

		go listener.ListenFlushCartEvents()

		ch <- &redis.Message{
			Channel: FlushCartChannelName,
			Payload: goodMessage,
		}

		close(ch)

		time.Sleep(100 * time.Millisecond)

		assert.Equal(t, 1, counter)
		assert.Equal(t, "434343", gotUserID)
	})
}
