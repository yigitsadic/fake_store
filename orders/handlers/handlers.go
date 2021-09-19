package handlers

import (
	"context"
	"github.com/yigitsadic/fake_store/orders/database"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
)

type Server struct {
	orders_grpc.UnimplementedOrdersServiceServer
	OrderRepository database.Repository
}

func (s *Server) ListOrders(_ context.Context, req *orders_grpc.OrderListRequest) (*orders_grpc.OrderListResponse, error) {
	orders, err := s.OrderRepository.FindAll(req.UserId)
	if err != nil {
		return nil, err
	}

	return &orders_grpc.OrderListResponse{Orders: orders.ConvertToGRPCModel()}, nil
}

func (s *Server) StartOrder(_ context.Context, req *orders_grpc.StartOrderRequest) (*orders_grpc.StartOrderResponse, error) {
	order, err := s.OrderRepository.Start(req.GetUserId(), convertGrpcCartItemsToProduct(req.GetCartItems()))
	if err != nil {
		return nil, err
	}

	return &orders_grpc.StartOrderResponse{
		Id:            order.ID,
		PaymentAmount: order.PaymentAmount,
	}, nil
}
