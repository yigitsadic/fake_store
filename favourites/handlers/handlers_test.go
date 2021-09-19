package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/fake_store/favourites/database"
	"github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"
	"testing"
)

func TestServer_ListFavourites(t *testing.T) {
	repo := &database.MockFavouriteProduct{
		Storage: map[string]*database.FavouriteProduct{
			"55": {
				ID:        "55",
				ProductID: "1231",
				UserID:    "444",
				Status:    favourites_grpc.Product_POPULATED,
				Title:     "Lorem ipsum",
				Image:     "image.png",
			},
		},
	}

	t.Run("it should list with user id", func(t *testing.T) {
		repo.ErrorOnFind = false

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.ListFavouritesRequest{UserID: "444"}
		res, err := s.ListFavourites(context.TODO(), req)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetProducts()))
	})

	t.Run("it should handle if anything goes wrong", func(t *testing.T) {
		repo.ErrorOnFind = true

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.ListFavouritesRequest{UserID: "444"}
		res, err := s.ListFavourites(context.TODO(), req)

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestServer_MarkFavourite(t *testing.T) {
	repo := &database.MockFavouriteProduct{
		Storage: map[string]*database.FavouriteProduct{},
	}

	t.Run("it should create favourite record", func(t *testing.T) {
		repo.ErrorOnMark = false

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "555",
			UserID:    "YYYY",
		}
		res, err := s.MarkFavourite(context.TODO(), req)

		assert.Nil(t, err)
		assert.True(t, res.GetSuccess())
	})

	t.Run("it should handle if anything goes wrong", func(t *testing.T) {
		repo.ErrorOnMark = true

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "555",
			UserID:    "YYYY",
		}
		res, err := s.MarkFavourite(context.TODO(), req)
		assert.NotNil(t, err)
		assert.False(t, res.GetSuccess())
	})
}

func TestServer_UnMarkFavourite(t *testing.T) {
	repo := &database.MockFavouriteProduct{
		Storage: map[string]*database.FavouriteProduct{},
	}

	t.Run("it should delete favourite record", func(t *testing.T) {
		repo.ErrorOnRevokeMark = false
		repo.Storage["55"] = &database.FavouriteProduct{
			ID:        "55",
			ProductID: "1231",
			UserID:    "444",
			Status:    favourites_grpc.Product_POPULATED,
			Title:     "Lorem ipsum",
			Image:     "image.png",
		}

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "1231",
			UserID:    "444",
		}

		res, err := s.UnMarkFavourite(context.TODO(), req)

		assert.Nil(t, err)
		assert.True(t, res.GetSuccess())
	})

	t.Run("it should handle if record not found", func(t *testing.T) {
		repo.ErrorOnRevokeMark = false

		delete(repo.Storage, "55")

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "1231",
			UserID:    "444",
		}

		res, err := s.UnMarkFavourite(context.TODO(), req)

		assert.NotNil(t, err)
		assert.False(t, res.GetSuccess())
	})

	t.Run("it should handle if anything goes wrong", func(t *testing.T) {
		repo.ErrorOnRevokeMark = true

		repo.Storage["55"] = &database.FavouriteProduct{
			ID:        "55",
			ProductID: "1231",
			UserID:    "444",
			Status:    favourites_grpc.Product_POPULATED,
			Title:     "Lorem ipsum",
			Image:     "image.png",
		}

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "1231",
			UserID:    "444",
		}

		res, err := s.UnMarkFavourite(context.TODO(), req)

		assert.NotNil(t, err)
		assert.False(t, res.GetSuccess())
	})
}

func TestServer_ProductInFavourite(t *testing.T) {
	repo := &database.MockFavouriteProduct{
		Storage: map[string]*database.FavouriteProduct{
			"55": {
				ID:        "55",
				ProductID: "1231",
				UserID:    "444",
				Status:    favourites_grpc.Product_POPULATED,
				Title:     "Lorem ipsum",
				Image:     "image.png",
			},
		},
	}

	t.Run("it should return true if found", func(t *testing.T) {
		repo.ErrorOnFindProduct = false

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "1231",
			UserID:    "444",
		}
		res, err := s.ProductInFavourite(context.TODO(), req)

		assert.Nil(t, err)
		assert.True(t, res.GetInFavourites())
	})

	t.Run("it should return false if error occurs", func(t *testing.T) {
		repo.ErrorOnFindProduct = true

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "1231",
			UserID:    "444",
		}
		res, err := s.ProductInFavourite(context.TODO(), req)

		delete(repo.Storage, "55")

		assert.Nil(t, err)
		assert.False(t, res.GetInFavourites())
	})

	t.Run("it should return false if not found", func(t *testing.T) {
		repo.ErrorOnFindProduct = false

		s := &Server{FavouriteRepository: repo}
		req := &favourites_grpc.FavouritesRequest{
			ProductID: "1251",
			UserID:    "500",
		}
		res, err := s.ProductInFavourite(context.TODO(), req)

		assert.Nil(t, err)
		assert.False(t, res.GetInFavourites())
	})
}
