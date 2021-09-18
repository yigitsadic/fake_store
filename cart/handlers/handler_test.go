package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
	"testing"
)

func TestServer_AddToCart(t *testing.T) {
	userID := "eeee"

	t.Run("it should add if cart not exist", func(t *testing.T) {
		mockRepo := initializeMockRepo(t, false, false, false)
		s := &Server{CartRepository: mockRepo}
		req := generateAddToCartRequest(t, userID)
		res, err := s.AddToCart(context.TODO(), req)

		require.Nil(t, err)
		assert.Equal(t, int32(1), res.GetItemCount())
		assert.Equal(t, req.GetProductId(), res.GetCartItems()[0].GetProductId())
		assert.Equal(t, req.GetTitle(), res.GetCartItems()[0].GetTitle())
		assert.Equal(t, req.GetDescription(), res.GetCartItems()[0].GetDescription())
		assert.Equal(t, req.GetPrice(), res.GetCartItems()[0].GetPrice())
		assert.Equal(t, req.GetImage(), res.GetCartItems()[0].GetImage())
		assert.True(t, res.GetCartItems()[0].GetId() != "")
		assert.Equal(t, mockRepo.Storage[userID].Items[0].ID, res.GetCartItems()[0].GetId())
	})

	t.Run("it should add if cart present", func(t *testing.T) {
		mockRepo := initializeMockRepo(t, false, false, false)
		mockRepo.Storage[userID] = &database.Cart{
			UserID: userID,
			Items: []*database.CartItem{
				{
					ID:          "eee",
					ProductID:   "eeeee",
					UserID:      userID,
					Title:       "eeee",
					Description: "eee",
					Image:       "eee",
					Price:       54.4,
				},
			},
		}

		s := &Server{CartRepository: mockRepo}
		req := generateAddToCartRequest(t, userID)
		res, err := s.AddToCart(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, int32(2), res.GetItemCount())
	})

	t.Run("it should return an error if something went wrong", func(t *testing.T) {
		mockRepo := initializeMockRepo(t, true, false, false)
		s := &Server{CartRepository: mockRepo}
		req := generateAddToCartRequest(t, userID)
		res, err := s.AddToCart(context.TODO(), req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestServer_CartContent(t *testing.T) {
	userID := "ee"
	mockRepo := initializeMockRepo(t, false, false, false)
	mockRepo.Storage[userID] = &database.Cart{
		UserID: userID,
		Items: []*database.CartItem{
			{
				ID:          "eee",
				ProductID:   "eeeee",
				UserID:      userID,
				Title:       "eeee",
				Description: "eee",
				Image:       "eee",
				Price:       54.4,
			},
		},
	}

	t.Run("it should respond correctly even cart doesn't exists", func(t *testing.T) {
		s := &Server{CartRepository: mockRepo}
		req := genenrateCartContentRequest(t, "NOT_EXIST")
		res, err := s.CartContent(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, int32(0), res.GetItemCount())
	})

	t.Run("it should respond correctly if cart exists", func(t *testing.T) {
		s := &Server{CartRepository: mockRepo}
		req := genenrateCartContentRequest(t, userID)
		res, err := s.CartContent(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, int32(1), res.GetItemCount())
	})

	t.Run("it should not return an error if something went wrong", func(t *testing.T) {
		badRepo := initializeMockRepo(t, false, true, false)
		s := &Server{CartRepository: badRepo}
		req := genenrateCartContentRequest(t, userID)
		res, err := s.CartContent(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, int32(0), res.GetItemCount())
	})
}

func initializeMockRepo(t *testing.T, errorOnAdd, errorOnDisplay, errorOnDelete bool) *database.MockCartRepository {
	t.Helper()

	return &database.MockCartRepository{
		Storage:        make(map[string]*database.Cart),
		ErrorOnAdd:     errorOnAdd,
		ErrorOnDelete:  errorOnDelete,
		ErrorOnDisplay: errorOnDisplay,
	}
}

func genenrateCartContentRequest(t *testing.T, userID string) *cart_grpc.CartContentRequest {
	t.Helper()

	return &cart_grpc.CartContentRequest{UserId: userID}
}

func generateAddToCartRequest(t *testing.T, userID string) *cart_grpc.AddToCartRequest {
	t.Helper()

	return &cart_grpc.AddToCartRequest{
		UserId:      userID,
		ProductId:   "eee",
		Title:       "eee",
		Description: "eee",
		Price:       52.23,
		Image:       "ee",
	}
}
