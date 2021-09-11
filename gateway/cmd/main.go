package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/yigitsadic/fake_store/auth/client/client"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yigitsadic/fake_store/gateway/graph"
	"github.com/yigitsadic/fake_store/gateway/graph/generated"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3035"
	}

	authConnection, authClient := acquireAuthConnection()
	defer authConnection.Close()

	productsConnection, productClient := acquireProductsConnection()
	defer productsConnection.Close()

	resolver := graph.Resolver{
		AuthClient:     authClient,
		ProductsClient: productClient,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	r := chi.NewRouter()
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	r.Use(middleware.Heartbeat("/readiness"))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("Server is up and running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func acquireAuthConnection() (*grpc.ClientConn, client.AuthServiceClient) {
	conn, err := grpc.Dial("auth:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Unable to acquire auth connection")
	}

	c := client.NewAuthServiceClient(conn)

	return conn, c
}

func acquireProductsConnection() (*grpc.ClientConn, product_grpc.ProductServiceClient) {
	conn, err := grpc.Dial("products:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Unable to acquire products connection")
	}

	c := product_grpc.NewProductServiceClient(conn)

	return conn, c
}
