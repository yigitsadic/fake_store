package database

import (
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

// Order struct represents order.
type Order struct {
	ID            string
	UserID        string
	CreatedAt     time.Time
	PaymentAmount float32
	Status        orders_grpc.Order_OrderStatus

	Products []Product
}

// ConvertToGRPCModel converts order to grpc compatible struct.
func (o Order) ConvertToGRPCModel() *orders_grpc.Order {
	var products []*orders_grpc.Product

	for _, product := range o.Products {
		products = append(products, product.ConvertToGRPCModel())
	}

	return &orders_grpc.Order{
		Id:            o.ID,
		UserId:        o.UserID,
		PaymentAmount: o.PaymentAmount,
		CreatedAt:     o.CreatedAt.UTC().Format(time.RFC3339),
		Status:        o.Status,
		Products:      products,
	}
}

// Product struct represents product.
type Product struct {
	ID          string
	Title       string
	Description string
	Image       string
	Price       float32
}

// ConvertToGRPCModel converts product to grpc compatible struct.
func (p Product) ConvertToGRPCModel() *orders_grpc.Product {
	return &orders_grpc.Product{
		Id:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Image:       p.Image,
	}
}

// CartItem represents item in cart.
type CartItem struct {
	ID          string
	ProductID   string
	Title       string
	Description string
	Image       string
	Price       float32
}

type Repository interface {
	FindAll(userID string) ([]*Order, error)
	StartOrder(userID string, cartItems []CartItem) (*Order, error)
	CompleteOrder(orderID string) error
}
