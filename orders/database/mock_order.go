package database

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

// MockOrderRepository will mock real database connection for testing purposes.
type MockOrderRepository struct {
	ErrorOnFindAll  bool
	ErrorOnStart    bool
	ErrorOnComplete bool

	Storage map[string]*Order

	CompleteCallCounter int
}

func (o *MockOrderRepository) FindAll(userID string) (OrderList, error) {
	if o.ErrorOnFindAll {
		return nil, errors.New("something went wrong")
	}

	var orders OrderList

	for _, order := range o.Storage {
		if order.UserID == userID && order.Status == orders_grpc.Order_COMPLETED {
			orders = append(orders, *order)
		}
	}

	return orders, nil
}

func (o *MockOrderRepository) Start(userID string, products []Product) (*Order, error) {
	if o.ErrorOnStart {
		return nil, errors.New("something went wrong")
	}

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

func (o *MockOrderRepository) Complete(orderID string) (string, error) {
	if o.ErrorOnComplete {
		return "", errors.New("unable to continue")
	}

	order, ok := o.Storage[orderID]
	if ok && order.Status != orders_grpc.Order_COMPLETED {
		o.CompleteCallCounter++

		order.Status = orders_grpc.Order_COMPLETED

		return order.UserID, nil
	} else {
		return "", errors.New("record not found")
	}
}
