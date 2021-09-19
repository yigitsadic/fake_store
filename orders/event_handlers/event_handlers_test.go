package event_handlers

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/orders/database"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"testing"
	"time"
)

func TestEventHandler_ListenPaymentCompleteEvents(t *testing.T) {
	repo := &database.MockOrderRepository{
		Storage: make(map[string]*database.Order),
	}
	repo.Storage["333"] = &database.Order{
		ID:            "333",
		UserID:        "4444",
		CreatedAt:     time.Time{},
		PaymentAmount: 0,
		Status:        orders_grpc.Order_STARTED,
		Products:      nil,
	}

	t.Run("it should handle bad message gracefully", func(t *testing.T) {
		repo.CompleteCallCounter = 0

		ch := make(chan *redis.Message)

		h := EventHandler{
			MessageChan:     ch,
			OrderRepository: repo,
		}

		go h.ListenPaymentCompleteEvents()

		ch <- &redis.Message{Channel: PaymentCompleteChannel, Payload: ""}

		close(ch)

		assert.Equal(t, 0, repo.CompleteCallCounter)
	})

	t.Run("it should handle when order not found", func(t *testing.T) {
		repo.CompleteCallCounter = 0

		ch := make(chan *redis.Message)

		h := EventHandler{
			MessageChan:     ch,
			OrderRepository: repo,
		}

		go h.ListenPaymentCompleteEvents()

		ch <- &redis.Message{Channel: PaymentCompleteChannel, Payload: `{"reference_id": "YGT"}`}

		close(ch)

		assert.Equal(t, 0, repo.CompleteCallCounter)
	})

	t.Run("it should publish message to flush cart if order completed", func(t *testing.T) {
		repo.CompleteCallCounter = 0

		ch := make(chan *redis.Message)
		calledOnce := false

		h := &EventHandler{
			MessageChan:     ch,
			OrderRepository: repo,
			FlushCartFunc: func(_ string) {
				calledOnce = true
			},
		}

		repo.CompleteCallCounter = 0

		go h.ListenPaymentCompleteEvents()

		ch <- &redis.Message{Channel: PaymentCompleteChannel, Payload: `{"reference_id": "333"}`}

		close(ch)

		for !calledOnce {
		}

		assert.True(t, calledOnce)
		assert.Equal(t, 1, repo.CompleteCallCounter)
	})
}
