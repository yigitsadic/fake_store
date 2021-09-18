package database

import (
	"errors"
	"github.com/bxcodec/faker/v3"
)

type MockCartRepository struct {
	ErrorOnAdd bool
	Storage    map[string]*Cart
}

func (c *MockCartRepository) FlushCart(userID string) {
	c.Storage[userID] = &Cart{UserID: userID}
}

func (c *MockCartRepository) FindCart(userID string) (*Cart, error) {
	r, ok := c.Storage[userID]
	if ok {
		return r, nil
	}

	return nil, errors.New("cart not found")
}

func (c *MockCartRepository) AddToCart(item *CartItem) error {
	if c.ErrorOnAdd {
		return errors.New("something went wrong")
	}

	item.ID = faker.UUIDHyphenated()

	cart, ok := c.Storage[item.UserID]
	if !ok {
		c.Storage[item.UserID] = &Cart{
			UserID: item.UserID,
			Items:  []*CartItem{item},
		}

		return nil
	}

	var items []*CartItem

	items = append(cart.Items, item)

	c.Storage[item.UserID] = &Cart{
		UserID: item.UserID,
		Items:  items,
	}

	return nil
}

func (c *MockCartRepository) RemoveFromCart(itemID, userID string) error {
	cart, ok := c.Storage[userID]
	if !ok {
		return errors.New("cart not present")
	}

	var present []*CartItem

	for _, item := range cart.Items {
		if item.ID != itemID {
			present = append(present, item)
		}
	}

	cart.Items = present

	return nil
}
