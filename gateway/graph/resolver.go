//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/yigitsadic/fake_store/auth/client/client"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

type Resolver struct {
	AuthClient     client.AuthServiceClient
	ProductsClient product_grpc.ProductServiceClient
	CartService    cart_grpc.CartServiceClient
}
