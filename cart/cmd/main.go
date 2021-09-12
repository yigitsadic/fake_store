package main

import (
	"fmt"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CartItem struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
}

var CartStorage map[string][]CartItem

func init() {
	CartStorage = make(map[string][]CartItem)
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	s := server{}

	cart_grpc.RegisterCartServiceServer(grpcServer, &s)

	log.Println("Started to serve cart grpc")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
