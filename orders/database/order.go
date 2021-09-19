package database

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

// OrderRepository will mock real database connection.
type OrderRepository struct {
	Storage map[string]*Order
}

func (o *OrderRepository) FindAll(userID string) (OrderList, error) {
	var orders OrderList

	for _, order := range o.Storage {
		if order.UserID == userID && order.Status == orders_grpc.Order_COMPLETED {
			orders = append(orders, *order)
		}
	}

	return orders, nil
}

func (o *OrderRepository) Start(userID string, products []Product) (*Order, error) {
	var total float32

	for _, product := range products {
		total += product.Price
	}

	order := &Order{
		ID:            faker.UUIDHyphenated(),
		UserID:        userID,
		CreatedAt:     time.Now().UTC(),
		PaymentAmount: total,
		Status:        orders_grpc.Order_STARTED,
		Products:      products,
	}

	o.Storage[order.ID] = order

	return order, nil
}

func (o *OrderRepository) Complete(orderID string) (string, error) {
	order, ok := o.Storage[orderID]
	if ok && order.Status != orders_grpc.Order_COMPLETED {
		order.Status = orders_grpc.Order_COMPLETED

		return order.UserID, nil
	} else {
		return "", errors.New("order not found")
	}
}
