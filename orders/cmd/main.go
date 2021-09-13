package main

import (
	"fmt"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	s := server{}

	orders_grpc.RegisterOrdersServiceServer(grpcServer, &s)

	log.Println("Started to serve order grpc")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
