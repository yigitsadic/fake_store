package database

import "errors"

// ProductRepository handles interactions with database for products.
type ProductRepository struct {
	Storage map[string]Product
}

// FetchOne fetches one record from database with matching ID, else returns error.
func (p *ProductRepository) FetchOne(ID string) (*Product, error) {
	product, ok := p.Storage[ID]
	if ok {
		return &product, nil
	}

	return nil, errors.New("product not found on database")
}

// FetchAll fetches all products from database.
func (p *ProductRepository) FetchAll() (products []Product) {
	for _, product := range p.Storage {
		products = append(products, product)
	}

	return
}

// NewProductRepo initializes ProductRepository with empty storage.
func NewProductRepo() *ProductRepository {
	return &ProductRepository{
		Storage: map[string]Product{},
	}
}
