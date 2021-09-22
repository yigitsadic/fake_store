package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
	"github.com/yigitsadic/fake_store/cart/events"
	"github.com/yigitsadic/fake_store/cart/handlers"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

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
	db := client.Database("fake_store")

	repo := &database.CartRepository{
		Storage: db,
	}

	// Start to listen messages from Redis.

	cartFlushChan := rdb.Subscribe(context.Background(), events.FlushCartChannelName)
	populateChan := rdb.Subscribe(context.Background(), events.CartItemPopulatedChannelName)

	ev := events.EventListener{
		Repository:               repo,
		FlushCartMessageChan:     cartFlushChan.Channel(),
		ProductInfoPopulatedChan: populateChan.Channel(),
	}

	go ev.ListenFlushCartEvents()
	go ev.ListenCartItemDataPopulationMessages()

	// Initiate gRPC server.

	grpcServer := grpc.NewServer()
	s := handlers.Server{
		CartRepository: repo,
		PublishPopulateFunc: func(cartItemID, productID string) {
			events.PublishToRedis(rdb, cartItemID, productID)
		},
	}

	cart_grpc.RegisterCartServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Println("Started to serve cart grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
