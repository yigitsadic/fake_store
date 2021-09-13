package main

import (
	"context"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

type server struct {
	orders_grpc.UnimplementedOrdersServiceServer
}

func (s *server) ListOrders(ctx context.Context, req *orders_grpc.OrderListRequest) (*orders_grpc.OrderListResponse, error) {
	products := []*orders_grpc.Product{
		{
			Id:          "825c2ca8-cfeb-4ba4-8b34-fb93f7958fa8",
			Title:       "Cornflakes",
			Price:       6.94,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "46541671-d9dd-4e99-9f40-c807e1b14f11",
			Title:       "Vaccum Bag - 14x20",
			Price:       4.97,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
	}

	o := orders_grpc.Order{
		PaymentAmount: 11.91,
		CreatedAt:     time.Now().UTC().Format(time.RFC3339),
		Products:      products,
	}

	return &orders_grpc.OrderListResponse{Orders: []*orders_grpc.Order{&o}}, nil
}
