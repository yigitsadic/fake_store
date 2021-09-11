package main

import (
	"context"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

type server struct {
	product_grpc.UnimplementedProductServiceServer
}

func (s *server) ListProducts(context.Context, *product_grpc.ProductListRequest) (*product_grpc.ProductList, error) {
	products := []*product_grpc.Product{
		{
			Id:          "825c2ca8-cfeb-4ba4-8b34-fb93f7958fa8",
			Title:       "Cornflakes",
			Price:       6.94,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "46541671-d9dd-4e99-9f40-c807e1b14f11",
			Title:       "Vaccum Bag - 14x20",
			Price:       4.97,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "c3af5841-4cfe-4ba0-874b-7c8ced576357",
			Title:       "Mustard - Dijon",
			Price:       1.25,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "966a9098-3252-4a43-9776-dd7f66e09d91",
			Title:       "Cheese - Le Cru Du Clocher",
			Price:       1.69,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "0fef08f2-cc56-4fd7-9137-b0ab561bc7a1",
			Title:       "Beef - Striploin",
			Price:       2.71,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "9f932b92-3433-4be2-8302-7ac4901c97d6",
			Title:       "Beef - Bones, Marrow",
			Price:       8.73,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "6bf9959e-cf2c-4039-9a31-30a9e90e8d7c",
			Title:       "V8 Pet",
			Price:       6.61,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "4f2a902a-446f-41da-9d12-521f9c83c94a",
			Title:       "Sauce - Fish 25 Ozf Bottle",
			Price:       2.49,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "3497030f-7239-4fce-bb73-f446e4fedc10",
			Title:       "Beef - Rib Eye Aaa",
			Price:       6.38,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "49d5f82e-d636-4d6c-8508-5429db7fd4b1",
			Title:       "Muffin Mix - Banana Nut",
			Price:       5.68,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "ee66f7e3-4bdd-4298-b790-43a2431c77ab",
			Title:       "Dawn Professionl Pot And Pan",
			Price:       4.89,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "992d3766-6022-4ee1-847e-f293f2488951",
			Title:       "Jameson - Irish Whiskey",
			Price:       1.12,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "95ca1986-9e39-485e-942e-927ac91aecde",
			Title:       "Bread Fig And Almond",
			Price:       2.58,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "eb46937c-12f5-4b9b-8ffa-7cf20871fbaf",
			Title:       "Vinegar - White",
			Price:       4.16,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
		{
			Id:          "c03020f2-fbf2-463d-9003-15e1901dc47a",
			Title:       "Bouq All Italian - Primerba",
			Price:       4.33,
			Description: "Lorem ipsum dolor sit amet",
			Image:       "https://via.placeholder.com/150",
		},
	}

	return &product_grpc.ProductList{Products: products}, nil
}
