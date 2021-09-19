package database

import (
	"errors"
	"github.com/bxcodec/faker/v3"
)

type FavouriteProductRepo struct {
	Storage map[string]*FavouriteProduct
}

func (m *FavouriteProductRepo) MarkFavourite(productID, userID string) error {
	record := &FavouriteProduct{
		ID:        faker.UUIDHyphenated(),
		ProductID: productID,
		UserID:    userID,
		Status:    0,
	}

	m.Storage[record.ID] = record

	return nil
}

func (m *FavouriteProductRepo) RevokeMarkFavourite(productID, userID string) error {
	var recordId *string

	for _, item := range m.Storage {
		if item.ProductID == productID && item.UserID == userID {
			recordId = &item.ID
		}
	}

	if recordId != nil {
		delete(m.Storage, *recordId)
	}

	return errors.New("record not found")
}

func (m *FavouriteProductRepo) FindFavourites(userID string) ([]FavouriteProduct, error) {
	var records []FavouriteProduct

	for _, item := range m.Storage {
		if item.UserID == userID {
			records = append(records, FavouriteProduct{
				ID:        item.ID,
				ProductID: item.ProductID,
				UserID:    item.UserID,
				Status:    item.Status,
				Title:     item.Title,
				Image:     item.Image,
			})
		}
	}

	return records, nil
}

func (m *FavouriteProductRepo) ProductInFavourite(productID, userID string) bool {
	for _, item := range m.Storage {
		if item.ProductID == productID && item.UserID == userID {
			return true
		}
	}

	return false
}
