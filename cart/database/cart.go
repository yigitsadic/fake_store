package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CartRepository connects to Mongo.
type CartRepository struct {
	Storage *mongo.Database
	Ctx     context.Context
}

// UpdateCartItem updates cart item's title, description, price and image.
// Cart item contains product ID at creation. Other fields populated from Redis pub/sub.
func (c *CartRepository) UpdateCartItem(message CartItemProductMessage) {
	c.Storage.Collection("cart").UpdateOne(c.Ctx,
		bson.M{
			"items.product_id": message.ProductID,
			"items._id":        message.CartItemID,
			"active":           true,
		},
		bson.M{
			"$set": bson.M{
				"items.$.title":       message.Title,
				"items.$.description": message.Description,
				"items.$.price":       message.Price,
				"items.$.image":       message.Image,
			},
		},
	)
}

// FlushCart removes all items in cart.
func (c *CartRepository) FlushCart(userID string) {
	c.Storage.Collection("cart").UpdateMany(c.Ctx,
		bson.M{"user_id": userID},
		bson.M{"$set": bson.M{"active": false}},
	)
}

// FindCart returns cart for given userID. If no document found, return empty cart.
func (c *CartRepository) FindCart(userID string) (*Cart, error) {
	var cart Cart

	err := c.Storage.Collection("cart").FindOne(c.Ctx, bson.M{"user_id": userID, "active": true}).Decode(&cart)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &Cart{Items: nil}, nil
		}

		return nil, err
	}

	return &cart, nil
}

// AddToCart adds product to cart.
func (c *CartRepository) AddToCart(userID string, productID string) (string, error) {
	item := CartItem{
		ID:        primitive.NewObjectID().Hex(),
		ProductID: productID,
	}

	res, err := c.Storage.Collection("cart").UpdateOne(c.Ctx,
		bson.M{"user_id": userID, "active": true},
		bson.M{
			"$push": bson.M{
				"items": item,
			},
		},
	)

	if err == nil && res.MatchedCount == 0 {
		var cart Cart
		cart.Active = true
		cart.UserID = userID
		cart.Items = []CartItem{item}

		_, err = c.Storage.Collection("cart").InsertOne(c.Ctx, cart)
		if err != nil {
			return "", err
		}

		return item.ID, nil
	}

	return "", err
}

// RemoveFromCart removes given item from cart.
func (c *CartRepository) RemoveFromCart(itemID, userID string) error {
	res, err := c.Storage.Collection("cart").UpdateOne(c.Ctx,
		bson.M{"user_id": userID},
		bson.M{
			"$pull": bson.M{
				"items": bson.M{"_id": itemID},
			},
		},
	)

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("no rows found")
	}

	return nil
}
