package database

import "github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"

type Cart struct {
	UserID string
	Items  []*CartItem
}

func (c *Cart) ConvertToGrpcModel() *cart_grpc.CartContentResponse {
	var items []*cart_grpc.CartItem

	for _, cartItem := range c.Items {
		items = append(items, cartItem.ConvertToGrpcModel())
	}

	return &cart_grpc.CartContentResponse{
		ItemCount: int32(len(c.Items)),
		CartItems: items,
	}
}

type CartItem struct {
	ID          string
	ProductID   string
	UserID      string
	Title       string
	Description string
	Image       string
	Price       float32
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
	AddToCart(item *CartItem) error
	RemoveFromCart(itemID, userID string) error
	FlushCart(userID string)
}
