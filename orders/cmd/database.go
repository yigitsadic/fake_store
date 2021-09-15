package main

import (
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"time"
)

type database map[string]*orders_grpc.Order

func newDatabase() database {
	return make(database)
}

func newSeededDatabase() database {
	db := newDatabase()

	products := []*orders_grpc.Product{
		{
			Id:          "825c2ca8-cfeb-4ba4-8b34-fb93f7958fa8",
			Title:       "Cornflakes",
			Price:       6.94,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "46541671-d9dd-4e99-9f40-c807e1b14f11",
			Title:       "Vaccum Bag - 14x20",
			Price:       4.97,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
	}

	o := orders_grpc.Order{
		Id:            "initial",
		PaymentAmount: 11.91,
		Status:        orders_grpc.Order_STARTED,
		CreatedAt:     time.Now().UTC().Format(time.RFC3339),
		Products:      products,
	}

	db[o.Id] = &o

	return db
}
