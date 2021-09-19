package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"reflect"
	"testing"
	"time"
)

func TestProduct_ConvertToGRPCModel(t *testing.T) {
	p := Product{
		ID:          "EEEE",
		Title:       "TTTTT",
		Description: "AAAAA",
		Image:       "CCCCC",
		Price:       56.23,
	}

	got := p.ConvertToGRPCModel()

	assert.Equal(t, reflect.TypeOf(got), reflect.TypeOf(&orders_grpc.Product{}))
	assert.Equal(t, p.ID, got.GetId())
	assert.Equal(t, p.Title, got.GetTitle())
	assert.Equal(t, p.Description, got.GetDescription())
	assert.Equal(t, p.Image, got.GetImage())
	assert.Equal(t, p.Price, got.GetPrice())
}

func TestOrder_ConvertToGRPCModel(t *testing.T) {
	p1 := Product{
		ID:          "EEEE",
		Title:       "TTTTT",
		Description: "AAAAA",
		Image:       "CCCCC",
		Price:       10,
	}

	p2 := Product{
		ID:          "EEEE",
		Title:       "TTTTT",
		Description: "AAAAA",
		Image:       "CCCCC",
		Price:       15.23,
	}

	o := Order{
		ID:            "CXSSSD",
		UserID:        "EEQE",
		CreatedAt:     time.Now(),
		PaymentAmount: p1.Price + p2.Price,
		Status:        orders_grpc.Order_COMPLETED,
		Products:      []Product{p1, p2},
	}

	got := o.ConvertToGRPCModel()

	assert.Equal(t, reflect.TypeOf(&orders_grpc.Order{}), reflect.TypeOf(got))
	assert.Equal(t, o.PaymentAmount, got.GetPaymentAmount())
	assert.Equal(t, o.ID, got.GetId())
	assert.Equal(t, o.CreatedAt.UTC().Format(time.RFC3339), got.GetCreatedAt())
	assert.Equal(t, o.UserID, got.GetUserId())
	assert.Equal(t, len(o.Products), len(got.GetProducts()))
}
