package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/orders/database"
	"github.com/yigitsadic/fake_store/orders/event_handlers"
	"github.com/yigitsadic/fake_store/orders/handlers"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	// Redis connection.
	rdb := redis.NewClient(
		&redis.Options{Addr: "redis:6379"},
	)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("Unable to connect redis")
	}

	// Mongo connection.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://database:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err)
	}

	// Inject Mongo client to repository.
	repo := &database.OrderRepository{
		Storage: client.Database("fake_store"),
		Ctx:     context.Background(),
	}

	// Start to listen messages from Redis.

	pubSub := rdb.Subscribe(context.Background(), event_handlers.PaymentCompleteChannel)

	eventHandler := event_handlers.EventHandler{
		FlushCartFunc: func(message string) {
			rdb.Publish(context.Background(), event_handlers.FlushCartChannel, message)
		},
		MessageChan:     pubSub.Channel(),
		OrderRepository: repo,
	}

	go eventHandler.ListenPaymentCompleteEvents()

	// Initiate gRPC server.

	grpcServer := grpc.NewServer()
	s := handlers.Server{OrderRepository: repo}

	orders_grpc.RegisterOrdersServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Println("Started to serve order grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
