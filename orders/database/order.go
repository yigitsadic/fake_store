package database

import (
	"context"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// OrderRepository handles operations in Mongo.
type OrderRepository struct {
	Storage *mongo.Database
	Ctx     context.Context
}

// FindAll fetches all completed order documents that matching with user id.
func (o *OrderRepository) FindAll(userID string) (OrderList, error) {
	cursor, err := o.Storage.Collection("orders").Find(o.Ctx, bson.M{
		"user_id": userID,
		"status":  orders_grpc.Order_COMPLETED,
	})
	if err != nil {
		return nil, err
	}

	var orders OrderList

	for cursor.Next(o.Ctx) {
		var order Order
		if err = cursor.Decode(&order); err == nil {
			orders = append(orders, order)
		}
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(o.Ctx)

	return orders, nil
}

// Start inserts new order record with state *STARTED*.
func (o *OrderRepository) Start(userID string, products []Product) (*Order, error) {
	var total float32

	for _, product := range products {
		total += product.Price
	}

	order := &Order{
		UserID:        userID,
		CreatedAt:     time.Now().UTC(),
		PaymentAmount: total,
		Status:        orders_grpc.Order_STARTED,
		Products:      products,
	}

	result, err := o.Storage.Collection("orders").InsertOne(o.Ctx, order)
	if err != nil {
		return nil, err
	}

	order.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return order, nil
}

// Complete updates status of given order id and returns userID.
func (o *OrderRepository) Complete(orderID string) (string, error) {
	var order Order

	err := o.Storage.Collection("orders").FindOne(o.Ctx, bson.M{"_id": orderID}).Decode(&order)
	if err != nil {
		return "", err
	}

	userId := order.UserID

	_, err = o.Storage.Collection("orders").UpdateOne(o.Ctx,
		bson.M{
			"_id": orderID,
		},
		bson.M{
			"status": orders_grpc.Order_COMPLETED,
		},
	)

	if err != nil {
		return "", err
	}

	return userId, nil
}
