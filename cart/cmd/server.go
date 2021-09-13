package main

import (
	"context"
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
)

type server struct {
	cart_grpc.UnimplementedCartServiceServer
}

func formatCartItemsToGrpcCompatible(items []CartItem) []*cart_grpc.CartItem {
	var buildItems []*cart_grpc.CartItem

	for _, item := range items {
		buildItems = append(buildItems, &cart_grpc.CartItem{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			Image:       item.Image,
		})
	}

	return buildItems
}
func (s *server) CartContent(ctx context.Context, req *cart_grpc.CartContentRequest) (*cart_grpc.CartContentResponse, error) {
	items, ok := CartStorage[req.GetUserId()]
	if ok {
		res := formatCartItemsToGrpcCompatible(items)

		return &cart_grpc.CartContentResponse{
			ItemCount: int32(len(res)),
			CartItems: res,
		}, nil
	} else {
		return &cart_grpc.CartContentResponse{
			ItemCount: 0,
			CartItems: nil,
		}, nil
	}
}

func (s *server) AddToCart(ctx context.Context, req *cart_grpc.AddToCartRequest) (*cart_grpc.CartContentResponse, error) {
	var formattedItems []*cart_grpc.CartItem
	items, ok := CartStorage[req.GetUserId()]
	cartItem := CartItem{
		ID:          faker.UUIDHyphenated(),
		ProductID:   req.GetProductId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Image:       req.GetImage(),
	}

	if ok {
		items = append(items, cartItem)
	} else {
		items = []CartItem{cartItem}
	}

	CartStorage[req.GetUserId()] = items
	formattedItems = formatCartItemsToGrpcCompatible(items)

	return &cart_grpc.CartContentResponse{
		ItemCount: int32(len(items)),
		CartItems: formattedItems,
	}, nil
}

func (s *server) RemoveFromCart(ctx context.Context, req *cart_grpc.RemoveFromCartRequest) (*cart_grpc.CartContentResponse, error) {
	items, ok := CartStorage[req.GetUserId()]
	if ok {
		var filteredItems []CartItem

		for _, item := range items {
			if item.ID != req.GetCartItemId() {
				filteredItems = append(filteredItems, item)
			}
		}

		CartStorage[req.GetUserId()] = filteredItems

		res := formatCartItemsToGrpcCompatible(items)

		return &cart_grpc.CartContentResponse{
			ItemCount: int32(len(res)),
			CartItems: res,
		}, nil
	} else {
		return nil, errors.New("no item found in cart")
	}
}
