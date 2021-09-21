package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/orders/database"
	"github.com/yigitsadic/fake_store/orders/event_handlers"
	"github.com/yigitsadic/fake_store/orders/handlers"
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

	rdb := redis.NewClient(
		&redis.Options{Addr: "redis:6379"},
	)

	if err = rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("Unable to connect redis")
	}

	repo := &database.OrderRepository{
		Storage: make(map[string]*database.Order),
	}

	pubSub := rdb.Subscribe(context.Background(), event_handlers.PaymentCompleteChannel)

	eventHandler := event_handlers.EventHandler{
		FlushCartFunc: func(message string) {
			rdb.Publish(context.Background(), event_handlers.FlushCartChannel, message)
		},
		MessageChan:     pubSub.Channel(),
		OrderRepository: repo,
	}

	go eventHandler.ListenPaymentCompleteEvents()

	grpcServer := grpc.NewServer()
	s := handlers.Server{OrderRepository: repo}

	orders_grpc.RegisterOrdersServiceServer(grpcServer, &s)

	log.Println("Started to serve order grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
