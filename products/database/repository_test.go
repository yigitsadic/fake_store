package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"reflect"
	"testing"
)

func TestProduct_ConvertToGRPC(t *testing.T) {
	p := Product{
		ID:          "ae",
		Title:       "eee",
		Description: "loo",
		Image:       "liii",
		Price:       17.5,
	}

	got := p.ConvertToGRPC()

	assert.Equalf(t, reflect.TypeOf(&product_grpc.Product{}), reflect.TypeOf(got), "expected to see product_grpc.Product")
	assert.Equal(t, p.ID, got.GetId())
	assert.Equal(t, p.Title, got.GetTitle())
	assert.Equal(t, p.Description, got.GetDescription())
	assert.Equal(t, p.Image, got.GetImage())
	assert.Equal(t, p.Price, got.GetPrice())
}
