package graph

import (
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/gateway/graph/model"
)

func ConvertCartFromService(cartItems []*cart_grpc.CartItem) []*model.CartItem {
	var buildItems []*model.CartItem

	for _, item := range cartItems {
		buildItems = append(buildItems, &model.CartItem{
			ID:          item.GetId(),
			ProductID:   item.GetProductId(),
			Title:       item.GetTitle(),
			Description: item.GetDescription(),
			Price:       float64(item.GetPrice()),
			Image:       item.GetImage(),
		})
	}

	return buildItems
}
