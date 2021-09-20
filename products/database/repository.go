package database

import "github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"

// Product is a struct for representing products.
type Product struct {
	ID          string  `bson:"_id,omitempty"`
	Title       string  `bson:"title,omitempty"`
	Description string  `bson:"description,omitempty"`
	Image       string  `bson:"image,omitempty"`
	Price       float32 `bson:"price,omitempty"`
}

// ConvertToGRPC converts Product struct to gRPC compatible struct.
func (p Product) ConvertToGRPC() *product_grpc.Product {
	return &product_grpc.Product{
		Id:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Image:       p.Image,
	}
}

// Repository is an interface to communicate with database.
type Repository interface {
	FetchAll() []Product
	FetchOne(string) (*Product, error)
}
