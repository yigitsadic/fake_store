package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestServer_AddToCart(t *testing.T) {
	userID := "eeee"

	t.Run("it should add if cart not exist", func(t *testing.T) {
		mockRepo := initializeMockRepo(t, false, false, false)
		s := &Server{CartRepository: mockRepo}
		req := generateAddToCartRequest(t, userID)
		_, err := s.AddToCart(context.TODO(), req)

		require.Nil(t, err)
	})

	t.Run("it should add if cart present", func(t *testing.T) {
		mockRepo := initializeMockRepo(t, false, false, false)
		mockRepo.Storage[userID] = &database.Cart{
			UserID: userID,
			Items: []database.CartItem{
				{
					ID:          primitive.NewObjectID().Hex(),
					ProductID:   "eeeee",
					Title:       "eeee",
					Description: "eee",
					Image:       "eee",
					Price:       54.4,
				},
			},
		}

		s := &Server{CartRepository: mockRepo}
		req := generateAddToCartRequest(t, userID)
		_, err := s.AddToCart(context.TODO(), req)

		assert.Nil(t, err)
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
		Items: []database.CartItem{
			{
				ID:          primitive.NewObjectID().Hex(),
				ProductID:   "eeeee",
				Title:       "eeee",
				Description: "eee",
				Image:       "eee",
				Price:       54.4,
			},
		},
	}

	t.Run("it should respond correctly even cart doesn't exists", func(t *testing.T) {
		s := &Server{CartRepository: mockRepo}
		req := generateCartContentRequest(t, "NOT_EXIST")
		res, err := s.CartContent(context.TODO(), req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("it should respond correctly if cart exists", func(t *testing.T) {
		s := &Server{CartRepository: mockRepo}
		req := generateCartContentRequest(t, userID)
		res, err := s.CartContent(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetCartItems()))
	})

	t.Run("it should not return an error if something went wrong", func(t *testing.T) {
		badRepo := initializeMockRepo(t, false, true, false)
		s := &Server{CartRepository: badRepo}
		req := generateCartContentRequest(t, userID)
		res, err := s.CartContent(context.TODO(), req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestServer_RemoveFromCart(t *testing.T) {
	userID := "eeee"
	cartItemID := primitive.NewObjectID().Hex()

	mockRepo := initializeMockRepo(t, false, false, false)
	mockRepo.Storage[userID] = &database.Cart{
		UserID: userID,
		Items: []database.CartItem{
			{
				ID:          cartItemID,
				ProductID:   "eeeee",
				Title:       "eeee",
				Description: "eee",
				Image:       "eee",
				Price:       54.4,
			},
		},
	}

	t.Run("it should return cart if not item found", func(t *testing.T) {
		s := &Server{CartRepository: mockRepo}
		req := generateRemoveFromCartRequest(t, primitive.NewObjectID().Hex(), userID)
		_, err := s.RemoveFromCart(context.TODO(), req)

		assert.Nil(t, err)
	})

	t.Run("it should return nil if everything went good", func(t *testing.T) {
		s := &Server{CartRepository: mockRepo}
		req := generateRemoveFromCartRequest(t, cartItemID, userID)
		_, err := s.RemoveFromCart(context.TODO(), req)

		assert.Nil(t, err)
	})

	t.Run("it should return an error if something went wrong", func(t *testing.T) {
		badRepo := initializeMockRepo(t, false, false, true)
		badRepo.Storage[userID] = &database.Cart{
			UserID: userID,
			Items: []database.CartItem{
				{
					ID:          cartItemID,
					ProductID:   "eeeee",
					Title:       "eeee",
					Description: "eee",
					Image:       "eee",
					Price:       54.4,
				},
			},
		}

		s := &Server{CartRepository: badRepo}
		req := generateRemoveFromCartRequest(t, cartItemID, userID)
		res, err := s.RemoveFromCart(context.TODO(), req)
		assert.Nil(t, res)
		assert.NotNil(t, err)
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

func generateCartContentRequest(t *testing.T, userID string) *cart_grpc.CartContentRequest {
	t.Helper()

	return &cart_grpc.CartContentRequest{UserId: userID}
}

func generateAddToCartRequest(t *testing.T, userID string) *cart_grpc.AddToCartRequest {
	t.Helper()

	return &cart_grpc.AddToCartRequest{
		UserId:    userID,
		ProductId: "eee",
	}
}

func generateRemoveFromCartRequest(t *testing.T, ID, userID string) *cart_grpc.RemoveFromCartRequest {
	t.Helper()

	return &cart_grpc.RemoveFromCartRequest{CartItemId: ID, UserId: userID}
}
