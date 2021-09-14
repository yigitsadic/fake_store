package main

import (
	"context"
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
)

type server struct {
	Database *cartDatabase
	cart_grpc.UnimplementedCartServiceServer
}

func (s *server) CartContent(_ context.Context, req *cart_grpc.CartContentRequest) (*cart_grpc.CartContentResponse, error) {
	items, ok := s.Database.Storage[req.GetUserId()]
	if ok {
		res := s.Database.formatCartItemsToGrpcCompatible(items)

		return &cart_grpc.CartContentResponse{
			ItemCount: int32(len(res)),
			CartItems: res,
		}, nil
	}

	return &cart_grpc.CartContentResponse{
		ItemCount: 0,
		CartItems: nil,
	}, nil
}

func (s *server) AddToCart(_ context.Context, req *cart_grpc.AddToCartRequest) (*cart_grpc.CartContentResponse, error) {
	var formattedItems []*cart_grpc.CartItem
	items, ok := s.Database.Storage[req.GetUserId()]
	item := cartItem{
		ID:          faker.UUIDHyphenated(),
		ProductID:   req.GetProductId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Image:       req.GetImage(),
	}

	if ok {
		items = append(items, item)
	} else {
		items = []cartItem{item}
	}

	s.Database.Storage[req.GetUserId()] = items
	formattedItems = s.Database.formatCartItemsToGrpcCompatible(items)

	return &cart_grpc.CartContentResponse{
		ItemCount: int32(len(items)),
		CartItems: formattedItems,
	}, nil
}

func (s *server) RemoveFromCart(_ context.Context, req *cart_grpc.RemoveFromCartRequest) (*cart_grpc.CartContentResponse, error) {
	items, ok := s.Database.Storage[req.GetUserId()]
	if ok {
		var filteredItems []cartItem

		for _, item := range items {
			if item.ID != req.GetCartItemId() {
				filteredItems = append(filteredItems, item)
			}
		}

		s.Database.Storage[req.GetUserId()] = filteredItems

		res := s.Database.formatCartItemsToGrpcCompatible(items)

		return &cart_grpc.CartContentResponse{
			ItemCount: int32(len(res)),
			CartItems: res,
		}, nil
	}

	return nil, errors.New("no item found in cart")
}
