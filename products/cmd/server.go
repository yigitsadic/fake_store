package main

import (
	"context"
	"errors"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

type server struct {
	product_grpc.UnimplementedProductServiceServer
}

func (s *server) ListProducts(context.Context, *product_grpc.ProductListRequest) (*product_grpc.ProductList, error) {
	return &product_grpc.ProductList{Products: productDB}, nil
}

func (s *server) ProductDetail(ctx context.Context, req *product_grpc.ProductDetailRequest) (*product_grpc.Product, error) {
	var found *product_grpc.Product

	for _, item := range productDB {
		if req.GetProductId() == item.GetId() {
			found = &product_grpc.Product{
				Id:          item.GetId(),
				Title:       item.GetTitle(),
				Description: item.GetDescription(),
				Price:       item.GetPrice(),
				Image:       item.GetImage(),
			}
		}
	}

	if found == nil {
		return nil, errors.New("product not found on products database")
	}

	return found, nil
}
