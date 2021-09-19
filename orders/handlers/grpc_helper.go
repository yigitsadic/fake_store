package handlers

import (
	"github.com/yigitsadic/fake_store/orders/database"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
)

// convertGrpcCartItemsToProduct converts grpc compatible structs to slice of Product struct
func convertGrpcCartItemsToProduct(items []*orders_grpc.CartItem) []database.Product {
	var products []database.Product

	for _, item := range items {
		products = append(products, database.Product{
			ID:          item.GetId(),
			Title:       item.GetTitle(),
			Description: item.GetDescription(),
			Image:       item.GetImage(),
			Price:       item.GetPrice(),
		})
	}

	return products
}
