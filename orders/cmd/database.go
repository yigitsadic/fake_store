package main

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

func convertProductFromGRPCModel(cartItem *orders_grpc.CartItem) Product {
	return Product{
		ID:          cartItem.GetId(),
		Title:       cartItem.GetTitle(),
		Description: cartItem.GetDescription(),
		Image:       cartItem.GetImage(),
		Price:       cartItem.GetPrice(),
	}
}

type database map[string]Order

func newDatabase() database {
	return make(database)
}
