package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/cart/database"
	"github.com/yigitsadic/fake_store/cart/event_listener"
	"github.com/yigitsadic/fake_store/cart/handlers"
	"google.golang.org/grpc"
	"log"
	"net"
)

type cartItem struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
}

type cartDatabase struct {
	Storage map[string][]cartItem
}

func newCartDatabase() *cartDatabase {
	return &cartDatabase{
		Storage: make(map[string][]cartItem),
	}
}

func (d *cartDatabase) formatCartItemsToGrpcCompatible(items []cartItem) []*cart_grpc.CartItem {
	var buildItems []*cart_grpc.CartItem

	for _, item := range items {
		buildItems = append(buildItems, &cart_grpc.CartItem{
			Id:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Price:       item.Price,
			Image:       item.Image,
		})
	}

	return buildItems
}

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

	repo := &database.CartRepository{
		Storage: make(map[string]*database.Cart),
	}

	grpcServer := grpc.NewServer()
	s := handlers.Server{CartRepository: repo}

	events := event_listener.EventListener{
		RedisClient: rdb,
		Ctx:         context.Background(),
		Repository:  repo,
	}

	go events.ListenFlushCartEvents()

	cart_grpc.RegisterCartServiceServer(grpcServer, &s)

	log.Println("Started to serve cart grpc")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
