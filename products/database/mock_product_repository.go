package database

import "errors"

// MockProductRepository handles interactions with database for products.
type MockProductRepository struct {
	Storage map[string]Product
}

// FetchOne fetches one record from database with matching ID, else returns error.
func (p *MockProductRepository) FetchOne(ID string) (*Product, error) {
	product, ok := p.Storage[ID]
	if ok {
		return &product, nil
	}

	return nil, errors.New("product not found on database")
}

// FetchAll fetches all products from database.
func (p *MockProductRepository) FetchAll() (products []Product) {
	for _, product := range p.Storage {
		products = append(products, product)
	}

	return
}

// NewProductRepo initializes MockProductRepository with empty storage.
func NewProductRepo() *MockProductRepository {
	return &MockProductRepository{
		Storage: map[string]Product{},
	}
}
