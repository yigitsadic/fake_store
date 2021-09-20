package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Storage *mongo.Database
	Ctx     context.Context
}

func (p *ProductRepository) FetchAll() []Product {
	c, err := p.Storage.Collection("product_catalog").Find(p.Ctx, bson.D{})
	if err != nil {
		return []Product{}
	}

	var products []Product

	c.All(p.Ctx, &products)

	return products
}

func (p *ProductRepository) FetchOne(s string) (*Product, error) {
	var product Product

	err := p.Storage.Collection("product_catalog").FindOne(p.Ctx, bson.D{{"_id", s}}).Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
