package database

import "github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"

type FavouriteProduct struct {
	ID        string
	ProductID string
	UserID    string
	Status    favourites_grpc.Product_ProductStatus

	Title string
	Image string
}

func (p FavouriteProduct) ConvertToGrpcStruct() *favourites_grpc.Product {
	return &favourites_grpc.Product{
		Id:        p.ID,
		ProductID: p.ProductID,
		UserID:    p.UserID,
		Status:    p.Status,
		Title:     p.Title,
		Image:     p.Image,
	}
}

type Repository interface {
	MarkFavourite(productID, userID string) error
	RevokeMarkFavourite(productID, userID string) error

	FindFavourites(userID string) ([]FavouriteProduct, error)
	ProductInFavourite(productID, userID string) bool
}
