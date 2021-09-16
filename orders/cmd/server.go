package main

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

type server struct {
	orders_grpc.UnimplementedOrdersServiceServer
	Database database
}

func (s *server) ListOrders(_ context.Context, req *orders_grpc.OrderListRequest) (*orders_grpc.OrderListResponse, error) {
	var orders []*orders_grpc.Order

	for _, order := range s.Database {
		if order.UserID == req.GetUserId() && order.Status == orders_grpc.Order_COMPLETED {
			orders = append(orders, order.ConvertToGRPCModel())
		}
	}

	return &orders_grpc.OrderListResponse{Orders: orders}, nil
}

func (s *server) StartOrder(_ context.Context, req *orders_grpc.StartOrderRequest) (*orders_grpc.StartOrderResponse, error) {
	id := faker.UUIDHyphenated()
	var total float32
	var products []Product

	for _, item := range req.GetCartItems() {
		total += item.GetPrice()

		products = append(products, convertProductFromGRPCModel(item))
	}

	res := &orders_grpc.StartOrderResponse{
		Id:            id,
		PaymentAmount: total,
	}

	s.Database[res.GetId()] = Order{
		ID:            res.GetId(),
		UserID:        req.GetUserId(),
		PaymentAmount: total,
		CreatedAt:     time.Now().UTC(),
		Status:        orders_grpc.Order_STARTED,
		Products:      products,
	}

	return res, nil
}
