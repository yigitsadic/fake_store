package main

import (
	"context"
	"errors"
	"github.com/yigitsadic/fake_store/products/database"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const (
	bufSize = 1024 * 1024
)

var (
	lis *bufconn.Listener

	testProduct = database.Product{
		ID:          "ABCDEF",
		Title:       "Test product",
		Description: "Test description",
		Image:       "test image",
		Price:       54.25,
	}

	errExpected = errors.New("product not found")
)

type mockProductRepo struct {
}

func (m mockProductRepo) FetchAll() []database.Product {
	return []database.Product{testProduct}
}

func (m mockProductRepo) FetchOne(s string) (*database.Product, error) {
	if s == testProduct.ID {
		return &testProduct, nil
	}

	return nil, errExpected
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	product_grpc.RegisterProductServiceServer(s, &server{
		Repository: &mockProductRepo{},
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestServer_ListProducts(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unexpected to get connection error. Err=%s", err)
	}

	defer conn.Close()

	c := product_grpc.NewProductServiceClient(conn)

	res, err := c.ListProducts(ctx, &product_grpc.ProductListRequest{})
	if err != nil {
		t.Errorf("unexpected to get an error for list products but got=%s", err)
	}

	products := res.GetProducts()

	if len(products) != 1 {
		t.Errorf("expected count was %d but got %d", 1, len(products))
	}
}

func TestServer_ProductDetail(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unexpected to get connection error. Err=%s", err)
	}

	defer conn.Close()

	c := product_grpc.NewProductServiceClient(conn)

	res, err := c.ProductDetail(ctx, &product_grpc.ProductDetailRequest{
		ProductId: testProduct.ID,
	})
	if err != nil {
		t.Errorf("unexpected to get an error for product detail but got=%s", err)
	}

	if res.GetId() != testProduct.ID {
		t.Errorf("expected product id was %s but got %s", testProduct.ID, res.GetId())
	}

	if res.GetTitle() != testProduct.Title {
		t.Errorf("expected product title was %s but got %s", testProduct.Title, res.GetTitle())
	}

	if res.GetDescription() != testProduct.Description {
		t.Errorf("expected product description was %s but got %s", testProduct.Description, res.GetDescription())
	}

	if res.GetImage() != testProduct.Image {
		t.Errorf("expected product image was %s but got %s", testProduct.Image, res.GetImage())
	}

	if res.GetPrice() != testProduct.Price {
		t.Errorf("expected product price was %f but got %f", testProduct.Price, res.GetPrice())
	}
}
