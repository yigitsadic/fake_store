package database

import (
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
)

type Cart struct {
	ID     string     `bson:"_id,omitempty"`
	UserID string     `bson:"user_id,omitempty"`
	Active bool       `bson:"active,omitempty"`
	Items  []CartItem `bson:"items,omitempty"`
}

func (c *Cart) ConvertToGrpcModel() *cart_grpc.CartContentResponse {
	var items []*cart_grpc.CartItem

	for _, cartItem := range c.Items {
		items = append(items, cartItem.ConvertToGrpcModel())
	}

	return &cart_grpc.CartContentResponse{
		CartItems: items,
	}
}

type CartItem struct {
	ID          string  `bson:"_id,omitempty"`
	ProductID   string  `bson:"product_id,omitempty"`
	Title       string  `bson:"title,omitempty"`
	Description string  `bson:"description,omitempty"`
	Image       string  `bson:"image,omitempty"`
	Price       float32 `bson:"price,omitempty"`
}

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

type Repository interface {
	FindCart(userID string) (*Cart, error)
	AddToCart(userID string, productID string) error
	RemoveFromCart(itemID, userID string) error
	FlushCart(userID string)
}
