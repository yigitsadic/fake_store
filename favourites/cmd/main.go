package main

import (
	"context"
	"fmt"
	"github.com/yigitsadic/fake_store/favourites/database"
	"github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"
	"github.com/yigitsadic/fake_store/favourites/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
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
	repo := &database.FavouriteProductRepo{
		Storage: client.Database("fake_store"),
		Ctx:     context.Background(),
	}

	// Initiate gRPC server.

	grpcServer := grpc.NewServer()
	server := &handlers.Server{
		FavouriteRepository: repo,
	}

	favourites_grpc.RegisterFavouritesServiceServer(grpcServer, server)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Println("Started to serve favourites grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
