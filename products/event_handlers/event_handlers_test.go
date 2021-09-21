package event_handlers

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/fake_store/products/database"
	"testing"
	"time"
)

func TestEventHandler_ListenProductPopulateMessages(t *testing.T) {
	repo := &database.MockProductRepository{
		Storage: map[string]database.Product{
			"populate": {
				ID:          "populate",
				Title:       "Lorem",
				Description: "pi",
				Image:       "image.png",
				Price:       42.45,
			},
		},
	}

	t.Run("it should handle bad message", func(t *testing.T) {
		callCounter := 0

		ch := make(chan *redis.Message)
		h := EventHandler{
			MessageChan:       ch,
			ProductRepository: repo,
			PopulateCartItemFunc: func(product database.Product) {
				callCounter++
			},
		}

		go h.ListenProductPopulateMessages()

		ch <- &redis.Message{Channel: ProductInfoPopulateChannel, Payload: ""}

		close(ch)

		time.Sleep(100 * time.Millisecond)

		assert.Equal(t, 0, callCounter)
	})

	t.Run("it should handle good message", func(t *testing.T) {
		callCounter := 0

		ch := make(chan *redis.Message)
		h := EventHandler{
			MessageChan:       ch,
			ProductRepository: repo,
			PopulateCartItemFunc: func(product database.Product) {
				callCounter++
			},
		}

		go h.ListenProductPopulateMessages()

		b, err := json.Marshal(populateProductRequestMessage{
			ProductID:  "populate",
			CartItemID: "QQQ",
		})

		require.Nil(t, err)

		payload := string(b)

		ch <- &redis.Message{Channel: ProductInfoPopulateChannel, Payload: payload}

		close(ch)

		time.Sleep(100 * time.Millisecond)

		assert.Equal(t, 1, callCounter)
	})

	t.Run("it should handle if product not found", func(t *testing.T) {
		callCounter := 0

		ch := make(chan *redis.Message)
		h := EventHandler{
			MessageChan:       ch,
			ProductRepository: repo,
			PopulateCartItemFunc: func(product database.Product) {
				callCounter++
			},
		}

		go h.ListenProductPopulateMessages()

		b, err := json.Marshal(populateProductRequestMessage{
			ProductID:  "YYY",
			CartItemID: "QQQ",
		})

		require.Nil(t, err)

		payload := string(b)

		ch <- &redis.Message{Channel: ProductInfoPopulateChannel, Payload: payload}

		close(ch)

		time.Sleep(100 * time.Millisecond)

		assert.Equal(t, 0, callCounter)
	})
}
