package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/orders/database"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"

	"testing"
)

func TestServer_ListOrders(t *testing.T) {
	mockRepo := &database.MockOrderRepository{Storage: map[string]*database.Order{}}
	mockRepo.Storage["aa"] = &database.Order{
		ID:            "eee",
		UserID:        "EE",
		CreatedAt:     time.Time{},
		PaymentAmount: 12,
		Status:        orders_grpc.Order_COMPLETED,
		Products:      nil,
	}

	t.Run("it should list orders successfully", func(t *testing.T) {
		mockRepo.ErrorOnFindAll = false

		s := &Server{OrderRepository: mockRepo}
		req := generateOrderListRequest(t, "EE")
		res, err := s.ListOrders(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetOrders()))
		assert.Equal(t, "eee", res.GetOrders()[0].GetId())
		assert.Equal(t, "EE", res.GetOrders()[0].GetUserId())
	})

	t.Run("it should return an error if something went wrong", func(t *testing.T) {
		mockRepo.ErrorOnFindAll = true

		s := &Server{OrderRepository: mockRepo}
		req := generateOrderListRequest(t, "EE")
		res, err := s.ListOrders(context.TODO(), req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("it should return empty list if nothing found", func(t *testing.T) {
		mockRepo.ErrorOnFindAll = false

		s := &Server{OrderRepository: mockRepo}
		req := generateOrderListRequest(t, "NonExistingUser")
		res, err := s.ListOrders(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, 0, len(res.GetOrders()))
	})
}

func TestServer_StartOrder(t *testing.T) {
	t.Run("it should start order successfully", func(t *testing.T) {
		assert.False(t, true)
	})

	t.Run("it should return an error if something went wrong", func(t *testing.T) {
		assert.False(t, true)
	})

	t.Run("it should return an error if cart item is empty", func(t *testing.T) {
		assert.False(t, true)
	})
}

func generateOrderListRequest(t *testing.T, userID string) *orders_grpc.OrderListRequest {
	t.Helper()

	return &orders_grpc.OrderListRequest{UserId: userID}
}
