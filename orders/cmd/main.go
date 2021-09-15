package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

var orderDatabase database

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("Unable to connect redis")
	}

	orderDatabase = newSeededDatabase()

	listener := eventListener{
		RedisClient: rdb,
		Ctx:         context.Background(),
		Database:    orderDatabase,
	}

	go listener.ListenPaymentCompleteEvents()

	grpcServer := grpc.NewServer()
	s := server{}

	orders_grpc.RegisterOrdersServiceServer(grpcServer, &s)

	log.Println("Started to serve order grpc")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
