package database

import (
	"errors"
	"github.com/bxcodec/faker/v3"
)

type MockFavouriteProduct struct {
	Storage map[string]*FavouriteProduct

	ErrorOnMark        bool
	ErrorOnRevokeMark  bool
	ErrorOnFind        bool
	ErrorOnFindProduct bool
}

func (m *MockFavouriteProduct) MarkFavourite(productID, userID string) error {
	if m.ErrorOnMark {
		return errors.New("something went wrong")
	}

	record := &FavouriteProduct{
		ID:        faker.UUIDHyphenated(),
		ProductID: productID,
		UserID:    userID,
		Status:    0,
	}

	m.Storage[record.ID] = record

	return nil
}

func (m *MockFavouriteProduct) RevokeMarkFavourite(productID, userID string) error {
	if m.ErrorOnRevokeMark {
		return errors.New("something went wrong")
	}

	var recordId *string

	for _, item := range m.Storage {
		if item.ProductID == productID && item.UserID == userID {
			recordId = &item.ID
		}
	}

	if recordId != nil {
		delete(m.Storage, *recordId)

		return nil
	}

	return errors.New("record not found")
}

func (m *MockFavouriteProduct) FindFavourites(userID string) ([]FavouriteProduct, error) {
	if m.ErrorOnFind {
		return nil, errors.New("something went wrong")
	}

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

func (m *MockFavouriteProduct) ProductInFavourite(productID, userID string) bool {
	if m.ErrorOnFindProduct {
		return false
	}

	for _, item := range m.Storage {
		if item.ProductID == productID && item.UserID == userID {
			return true
		}
	}

	return false
}
