package database

import "github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"

// FavouriteProduct represents product in favourite data.
type FavouriteProduct struct {
	ID        string                                `bson:"_id,omitempty"`
	ProductID string                                `bson:"product_id,omitempty"`
	UserID    string                                `bson:"user_id,omitempty"`
	Status    favourites_grpc.Product_ProductStatus `bson:"status,omitempty"`

	Title string `bson:"title,omitempty"`
	Image string `bson:"image,omitempty"`
}

// ConvertToGrpcStruct converts struct to gRPC compatible struct.
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

// Repository represents interface for required functionalities that handler requires.
type Repository interface {
	MarkFavourite(productID, userID string) error
	RevokeMarkFavourite(productID, userID string) error

	FindFavourites(userID string) ([]FavouriteProduct, error)
	ProductInFavourite(productID, userID string) bool
}
