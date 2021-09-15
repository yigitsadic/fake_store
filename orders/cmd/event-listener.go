package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

type paymentIntentMessage struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type flushCartMessage struct {
	UserID string `json:"user_id"`
}

type eventListener struct {
	RedisClient *redis.Client
	Ctx         context.Context
	Database    database
}

func (l *eventListener) ListenPaymentCompleteEvents() {
	pubSub := l.RedisClient.Subscribe(l.Ctx, "PAYMENTS_COMPLETE_CHANNEL")

	ch := pubSub.Channel()

	for msg := range ch {
		var paymentMessage paymentIntentMessage

		err := json.Unmarshal([]byte(msg.Payload), &paymentMessage)
		if err == nil {
			record, ok := l.Database[paymentMessage.ReferenceID]
			if ok {
				newRecord := orders_grpc.Order{
					Id:            record.GetId(),
					UserId:        record.GetUserId(),
					PaymentAmount: record.GetPaymentAmount(),
					CreatedAt:     record.GetCreatedAt(),
					Products:      record.GetProducts(),
					Status:        orders_grpc.Order_COMPLETED,
				}

				l.Database[paymentMessage.ReferenceID] = &newRecord

				b, err := json.Marshal(flushCartMessage{UserID: newRecord.GetUserId()})
				if err == nil {
					l.RedisClient.Publish(l.Ctx, "FLUSH_CART_CHANNEL", string(b))
				}
			}
		}
	}
}
