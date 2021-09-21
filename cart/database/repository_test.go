package database

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func TestCartItem_ConvertToGrpcModel(t *testing.T) {
	item := CartItem{
		ID:          primitive.NewObjectID().Hex(),
		ProductID:   faker.UUIDHyphenated(),
		UserID:      faker.UUIDHyphenated(),
		Title:       faker.UUIDHyphenated(),
		Description: faker.UUIDHyphenated(),
		Image:       faker.UUIDHyphenated(),
		Price:       15.25,
	}

	got := item.ConvertToGrpcModel()

	assert.Equal(t, reflect.TypeOf(got), reflect.TypeOf(&cart_grpc.CartItem{}))
	assert.Equal(t, item.ID, got.GetId())
	assert.Equal(t, item.ProductID, got.GetProductId())
	assert.Equal(t, item.Title, got.GetTitle())
	assert.Equal(t, item.Description, got.GetDescription())
	assert.Equal(t, item.Image, got.GetImage())
	assert.Equal(t, item.Price, got.GetPrice())
}

func TestCart_ConvertToGrpcModel(t *testing.T) {
	cart := Cart{
		UserID: faker.UUIDHyphenated(),
		Items: []CartItem{
			{
				ID:          primitive.NewObjectID().Hex(),
				ProductID:   faker.UUIDHyphenated(),
				UserID:      faker.UUIDHyphenated(),
				Title:       faker.UUIDHyphenated(),
				Description: faker.UUIDHyphenated(),
				Image:       faker.UUIDHyphenated(),
				Price:       15.25,
			},
		},
	}

	got := cart.ConvertToGrpcModel()
	item := cart.Items[0]
	gotCartItem := got.GetCartItems()[0]

	assert.Equal(t, len(cart.Items), len(got.GetCartItems()))

	assert.Equal(t, reflect.TypeOf(got), reflect.TypeOf(&cart_grpc.CartContentResponse{}))
	assert.Equal(t, item.ID, gotCartItem.GetId())
	assert.Equal(t, item.ProductID, gotCartItem.GetProductId())
	assert.Equal(t, item.Title, gotCartItem.GetTitle())
	assert.Equal(t, item.Description, gotCartItem.GetDescription())
	assert.Equal(t, item.Image, gotCartItem.GetImage())
	assert.Equal(t, item.Price, gotCartItem.GetPrice())
}
