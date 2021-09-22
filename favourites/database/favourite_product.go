package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FavouriteProductRepo queries and mutates mongodb collection.
type FavouriteProductRepo struct {
	Storage *mongo.Database
	Ctx     context.Context
}

// MarkFavourite marks given product in favourites of given user id.
func (m *FavouriteProductRepo) MarkFavourite(productID, userID string) error {
	record := FavouriteProduct{
		ProductID: productID,
		UserID:    userID,
	}

	_, err := m.Storage.Collection("favourites").InsertOne(
		m.Ctx,
		record,
	)
	if err != nil {
		return err
	}

	return nil
}

// RevokeMarkFavourite deletes user-product favourite records.
func (m *FavouriteProductRepo) RevokeMarkFavourite(productID, userID string) error {
	_, err := m.Storage.Collection("favourites").DeleteMany(m.Ctx, bson.M{"product_id": productID, "user_id": userID})
	if err != nil {
		return err
	}

	return nil
}

// FindFavourites lists given user's favourite products.
func (m *FavouriteProductRepo) FindFavourites(userID string) ([]FavouriteProduct, error) {
	var records []FavouriteProduct

	cur, err := m.Storage.Collection("favourites").Find(m.Ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	for cur.Next(m.Ctx) {
		var record FavouriteProduct

		if err = cur.Decode(&record); err != nil {
			continue
		}

		records = append(records, record)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(m.Ctx)

	return records, nil
}

// ProductInFavourite queries does given product in user's favourites list.
func (m *FavouriteProductRepo) ProductInFavourite(productID, userID string) bool {
	var res FavouriteProduct

	err := m.Storage.Collection("favourites").FindOne(
		m.Ctx,
		bson.M{"user_id": userID, "product_id": productID},
	).Decode(&res)

	if err != nil {
		return false
	}

	return res.ID != ""
}
