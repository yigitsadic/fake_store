//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/yigitsadic/fake_store/auth/auth_grpc/auth_grpc"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

// Resolver Dependency injection struct for schema resolvers.
type Resolver struct {
	AuthService     auth_grpc.AuthServiceClient
	ProductsService product_grpc.ProductServiceClient
	CartService     cart_grpc.CartServiceClient
	OrdersService   orders_grpc.OrdersServiceClient
}
