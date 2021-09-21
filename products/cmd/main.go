package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/products/database"
	"github.com/yigitsadic/fake_store/products/event_handlers"
	"github.com/yigitsadic/fake_store/products/handlers"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"go.mongodb.org/mongo-driver/bson"
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

	rdb := redis.NewClient(
		&redis.Options{Addr: "redis:6379"},
	)

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

	repo := &database.ProductRepository{Storage: db}

	pubSub := rdb.Subscribe(context.Background(), event_handlers.ProductInfoPopulateChannel)
	eventHandler := event_handlers.EventHandler{
		PopulateCartItemFunc: func(product database.Product) {
			b, err := json.Marshal(product)
			if err != nil {
				return
			}

			rdb.Publish(context.Background(), event_handlers.ProductPopulatedChannel, string(b))
		},
		MessageChan:       pubSub.Channel(),
		ProductRepository: repo,
	}

	go eventHandler.ListenProductPopulateMessages()

	count, err := db.Collection("product_catalog").CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatalln(err)
	}

	if count == 0 {
		var list []interface{}

		products := database.SeedDatabase()

		for _, item := range products {
			list = append(list, item)
		}

		db.Collection("product_catalog").InsertMany(ctx, list)
	}

	grpcServer := grpc.NewServer()
	s := handlers.Server{
		Repository: repo,
	}

	product_grpc.RegisterProductServiceServer(grpcServer, &s)

	log.Println("Started to serve product grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
