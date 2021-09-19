package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"testing"
)

func Test_convertGrpcCartItemsToProduct(t *testing.T) {
	given := []*orders_grpc.CartItem{
		{
			Id:          "12",
			ProductId:   "eee",
			Title:       "eee",
			Description: "eee",
			Price:       12.5,
			Image:       "eee",
		},
		{
			Id:          "13",
			ProductId:   "ddd",
			Title:       "ddd",
			Description: "ddd",
			Price:       24.5,
			Image:       "ddd",
		},
	}

	got := convertGrpcCartItemsToProduct(given)

	assert.Equal(t, len(given), len(got))

	for i, item := range got {
		assert.Equal(t, item.ID, given[i].GetId())
		assert.Equal(t, item.Title, given[i].GetTitle())
		assert.Equal(t, item.Description, given[i].GetDescription())
		assert.Equal(t, item.Price, given[i].GetPrice())
		assert.Equal(t, item.Image, given[i].GetImage())
	}
}
