package main

import (
	"fmt"
	"github.com/yigitsadic/fake_store/favourites/database"
	"github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"
	"github.com/yigitsadic/fake_store/favourites/handlers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	repo := &database.FavouriteProductRepo{
		Storage: map[string]*database.FavouriteProduct{},
	}
	grpcServer := grpc.NewServer()
	server := &handlers.Server{
		FavouriteRepository: repo,
	}

	favourites_grpc.RegisterFavouritesServiceServer(grpcServer, server)

	log.Println("Started to serve favourites grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
