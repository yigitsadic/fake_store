package database

import (
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
)

// CartItemProductMessage struct represents update cart function call parameter.
type CartItemProductMessage struct {
	ProductID   string  `json:"product_id"`
	CartItemID  string  `json:"cart_item_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float32 `json:"price"`
}

// Cart struct represents cart document in database.
type Cart struct {
	ID     string     `bson:"_id,omitempty"`
	UserID string     `bson:"user_id,omitempty"`
	Active bool       `bson:"active,omitempty"`
	Items  []CartItem `bson:"items,omitempty"`
}

// ConvertToGrpcModel converts Cart struct to gRPC compatible struct.
func (c *Cart) ConvertToGrpcModel() *cart_grpc.CartContentResponse {
	var items []*cart_grpc.CartItem

	for _, cartItem := range c.Items {
		items = append(items, cartItem.ConvertToGrpcModel())
	}

	return &cart_grpc.CartContentResponse{
		CartItems: items,
	}
}

// CartItem struct represents cart items in a cart.
type CartItem struct {
	ID          string  `bson:"_id,omitempty"`
	ProductID   string  `bson:"product_id,omitempty"`
	Title       string  `bson:"title,omitempty"`
	Description string  `bson:"description,omitempty"`
	Image       string  `bson:"image,omitempty"`
	Price       float32 `bson:"price,omitempty"`
}

// ConvertToGrpcModel converts cart item to gRPC compatible struct.
func (c *CartItem) ConvertToGrpcModel() *cart_grpc.CartItem {
	return &cart_grpc.CartItem{
		Id:          c.ID,
		ProductId:   c.ProductID,
		Title:       c.Title,
		Description: c.Description,
		Price:       c.Price,
		Image:       c.Image,
	}
}

// Repository is an interface that contains required functionalities by handlers and event handlers.
type Repository interface {
	FindCart(userID string) (*Cart, error)
	AddToCart(userID string, productID string) (string, error)
	RemoveFromCart(itemID, userID string) error
	FlushCart(userID string)
	UpdateCartItem(message CartItemProductMessage)
}
