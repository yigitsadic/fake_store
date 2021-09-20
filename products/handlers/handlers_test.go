package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/products/database"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"testing"
)

func TestServer_ListProducts(t *testing.T) {
	mockRepo := &database.MockProductRepository{Storage: map[string]database.Product{}}
	mockRepo.Storage["aa"] = database.Product{
		ID:          "eee",
		Title:       "EE",
		Description: "Hello there",
		Price:       12,
		Image:       "adqd.png",
	}

	s := &Server{Repository: mockRepo}

	res, err := s.ListProducts(context.TODO(), &product_grpc.ProductListRequest{})

	assert.Nil(t, err)

	products := res.GetProducts()
	assert.Equalf(t, 1, len(products), "expected count was %d but got %d", 1, len(products))
}

func TestServer_ProductDetail(t *testing.T) {
	testProduct := database.Product{
		ID:          "eee",
		Title:       "EE",
		Description: "Hello there",
		Price:       12,
		Image:       "adqd.png",
	}

	mockRepo := &database.MockProductRepository{Storage: map[string]database.Product{}}
	mockRepo.Storage[testProduct.ID] = testProduct

	s := &Server{Repository: mockRepo}

	res, err := s.ProductDetail(context.TODO(), &product_grpc.ProductDetailRequest{
		ProductId: testProduct.ID,
	})

	assert.Nilf(t, err, "unexpected to get an error for product detail but got=%s", err)
	assert.Equalf(t, testProduct.ID, res.GetId(), "expected product id was %s but got %s", testProduct.ID, res.GetId())
	assert.Equalf(t, testProduct.Title, res.GetTitle(), "expected product title was %s but got %s", testProduct.Title, res.GetTitle())
	assert.Equalf(t, testProduct.Description, res.GetDescription(), "expected product description was %s but got %s", testProduct.Description, res.GetDescription())
	assert.Equalf(t, testProduct.Image, res.GetImage(), "expected product image was %s but got %s", testProduct.Image, res.GetImage())
	assert.Equalf(t, testProduct.Price, res.GetPrice(), "expected product price was %f but got %f", testProduct.Price, res.GetPrice())
}
