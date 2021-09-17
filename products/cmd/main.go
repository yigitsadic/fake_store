package main

import (
	"fmt"
	"github.com/yigitsadic/fake_store/products/database"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	repo := database.NewProductRepo()

	database.SeedDatabase(repo)

	grpcServer := grpc.NewServer()
	s := server{
		Repository: repo,
	}

	product_grpc.RegisterProductServiceServer(grpcServer, &s)

	log.Println("Started to serve product grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
