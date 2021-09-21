package database

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockCartRepository struct {
	ErrorOnAdd     bool
	ErrorOnDisplay bool
	ErrorOnDelete  bool

	Storage map[string]*Cart
}

func (c *MockCartRepository) FlushCart(userID string) {
	c.Storage[userID] = &Cart{UserID: userID}
}

func (c *MockCartRepository) FindCart(userID string) (*Cart, error) {
	if c.ErrorOnDisplay {
		return nil, errors.New("not found")
	}

	r, ok := c.Storage[userID]
	if ok {
		return r, nil
	}

	return nil, errors.New("cart not found")
}

func (c *MockCartRepository) AddToCart(userID string, productID string) error {
	if c.ErrorOnAdd {
		return errors.New("something went wrong")
	}

	item := CartItem{
		ID:        primitive.NewObjectID().Hex(),
		ProductID: productID,
	}

	cart, ok := c.Storage[userID]
	if !ok {
		c.Storage[userID] = &Cart{
			UserID: userID,
			Items:  []CartItem{item},
		}

		return nil
	}

	var items []CartItem

	items = append(cart.Items, item)

	c.Storage[userID] = &Cart{
		UserID: userID,
		Items:  items,
	}

	return nil
}

func (c *MockCartRepository) RemoveFromCart(itemID, userID string) error {
	if c.ErrorOnDelete {
		return errors.New("unable to remove")
	}

	cart, ok := c.Storage[userID]
	if !ok {
		return errors.New("cart not present")
	}

	var present []CartItem

	for _, item := range cart.Items {
		if item.ID != itemID {
			present = append(present, item)
		}
	}

	cart.Items = present

	return nil
}
