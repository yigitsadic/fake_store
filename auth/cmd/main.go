package main

import (
	"fmt"
	"github.com/yigitsadic/fake_store/auth/auth_grpc/auth_grpc"
	"github.com/yigitsadic/fake_store/auth/handlers"
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
	s := handlers.Server{}

	auth_grpc.RegisterAuthServiceServer(grpcServer, &s)

	log.Println("Started to serve auth grpc")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
