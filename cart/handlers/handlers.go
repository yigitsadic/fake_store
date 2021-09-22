package handlers

import (
	"context"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
)

// Server is a struct that works as dependency injection to gRPC server functions.
type Server struct {
	CartRepository      database.Repository
	PublishPopulateFunc func(cartItemID, productID string)

	cart_grpc.UnimplementedCartServiceServer
}

// CartContent returns cart content. If nothing found on database will return empty cart as response.
func (s *Server) CartContent(_ context.Context, req *cart_grpc.CartContentRequest) (*cart_grpc.CartContentResponse, error) {
	cart, err := s.CartRepository.FindCart(req.GetUserId())
	if err != nil {
		return nil, err
	}

	return cart.ConvertToGrpcModel(), nil
}

// AddToCart adds given productID to cart. If no cart found, it will create cart and insert given product.
// Request contains only productId, because of that product data needed. For populating product data,
// a message will be fired to Redis pub/sub.
func (s *Server) AddToCart(_ context.Context, req *cart_grpc.AddToCartRequest) (*cart_grpc.CartOperation, error) {
	itemID, err := s.CartRepository.AddToCart(req.GetUserId(), req.GetProductId())

	if err != nil {
		return nil, err
	}

	go s.PublishPopulateFunc(itemID, req.GetProductId())

	return &cart_grpc.CartOperation{}, nil
}

// RemoveFromCart removes given cart item from cart.
func (s *Server) RemoveFromCart(_ context.Context, req *cart_grpc.RemoveFromCartRequest) (*cart_grpc.CartOperation, error) {
	err := s.CartRepository.RemoveFromCart(req.GetCartItemId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &cart_grpc.CartOperation{}, nil
}
