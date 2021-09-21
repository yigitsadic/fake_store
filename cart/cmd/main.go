package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
	"github.com/yigitsadic/fake_store/cart/event_listener"
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if err = rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("Unable to connect redis")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://database:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err)
	}

	db := client.Database("fake_store")

	repo := &database.CartRepository{
		Storage: db,
	}

	grpcServer := grpc.NewServer()
	s := handlers.Server{CartRepository: repo}

	pubSub := rdb.Subscribe(context.Background(), event_listener.ChannelName)

	events := event_listener.EventListener{
		Repository:  repo,
		MessageChan: pubSub.Channel(),
	}

	go events.ListenFlushCartEvents()

	cart_grpc.RegisterCartServiceServer(grpcServer, &s)

	log.Println("Started to serve cart grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
