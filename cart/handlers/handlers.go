package handlers

import (
	"context"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
)

type Server struct {
	CartRepository database.Repository
	cart_grpc.UnimplementedCartServiceServer
}

func (s *Server) CartContent(_ context.Context, req *cart_grpc.CartContentRequest) (*cart_grpc.CartContentResponse, error) {
	cart, err := s.CartRepository.FindCart(req.GetUserId())
	if err != nil {
		return nil, err
	}

	return cart.ConvertToGrpcModel(), nil
}

func (s *Server) AddToCart(_ context.Context, req *cart_grpc.AddToCartRequest) (*cart_grpc.CartOperation, error) {
	if err := s.CartRepository.AddToCart(req.GetUserId(), req.GetProductId()); err != nil {
		return nil, err
	}

	return &cart_grpc.CartOperation{}, nil
}

func (s *Server) RemoveFromCart(_ context.Context, req *cart_grpc.RemoveFromCartRequest) (*cart_grpc.CartOperation, error) {
	err := s.CartRepository.RemoveFromCart(req.GetCartItemId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &cart_grpc.CartOperation{}, nil
}
