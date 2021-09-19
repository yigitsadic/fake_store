package handlers

import (
	"context"
	"github.com/yigitsadic/fake_store/favourites/database"
	"github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"
)

type Server struct {
	FavouriteRepository database.Repository
	favourites_grpc.UnimplementedFavouritesServiceServer
}

func (s *Server) ListFavourites(ctx context.Context, request *favourites_grpc.ListFavouritesRequest) (*favourites_grpc.ListFavouritesResponse, error) {
	result, err := s.FavouriteRepository.FindFavourites(request.GetUserID())
	if err != nil {
		return nil, err
	}

	var products []*favourites_grpc.Product

	for _, item := range result {
		products = append(products, item.ConvertToGrpcStruct())
	}

	return &favourites_grpc.ListFavouritesResponse{
		Products: products,
	}, nil
}

func (s *Server) MarkFavourite(ctx context.Context, request *favourites_grpc.FavouritesRequest) (*favourites_grpc.FavouritesResponse, error) {
	err := s.FavouriteRepository.MarkFavourite(request.GetProductID(), request.GetUserID())

	if err != nil {
		return nil, err
	}

	return &favourites_grpc.FavouritesResponse{Success: true}, nil
}

func (s *Server) UnMarkFavourite(ctx context.Context, request *favourites_grpc.FavouritesRequest) (*favourites_grpc.FavouritesResponse, error) {
	err := s.FavouriteRepository.RevokeMarkFavourite(request.GetProductID(), request.GetUserID())

	if err != nil {
		return nil, err
	}

	return &favourites_grpc.FavouritesResponse{Success: true}, nil
}

func (s *Server) ProductInFavourite(ctx context.Context, request *favourites_grpc.FavouritesRequest) (*favourites_grpc.ProductInFavouritesResponse, error) {
	val := s.FavouriteRepository.ProductInFavourite(request.GetProductID(), request.GetUserID())

	return &favourites_grpc.ProductInFavouritesResponse{InFavourites: val}, nil
}
