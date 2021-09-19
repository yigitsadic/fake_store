package handlers

import (
	"context"
	"github.com/yigitsadic/fake_store/products/database"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

type Server struct {
	product_grpc.UnimplementedProductServiceServer

	Repository database.Repository
}

func (s *Server) ListProducts(context.Context, *product_grpc.ProductListRequest) (*product_grpc.ProductList, error) {
	var products []*product_grpc.Product

	for _, product := range s.Repository.FetchAll() {
		products = append(products, product.ConvertToGRPC())
	}

	return &product_grpc.ProductList{Products: products}, nil
}

func (s *Server) ProductDetail(ctx context.Context, req *product_grpc.ProductDetailRequest) (*product_grpc.Product, error) {
	product, err := s.Repository.FetchOne(req.GetProductId())
	if err != nil {
		return nil, err
	}

	return product.ConvertToGRPC(), nil
}
