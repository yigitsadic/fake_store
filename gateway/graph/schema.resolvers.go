package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/auth/client/client"
	"github.com/yigitsadic/fake_store/gateway/graph/generated"
	"github.com/yigitsadic/fake_store/gateway/graph/model"
	"github.com/yigitsadic/fake_store/gateway/helper"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
	"log"
)

var CartItems []*model.CartItem

func (r *mutationResolver) Login(ctx context.Context) (*model.LoginResponse, error) {
	result, err := r.AuthClient.LoginUser(ctx, &client.AuthRequest{})
	if err != nil {
		return nil, err
	}

	res := model.LoginResponse{
		ID:       result.Id,
		Avatar:   result.Avatar,
		FullName: result.FullName,
		Token:    result.JwtToken,
	}

	return &res, nil
}

func (r *mutationResolver) AddToCart(ctx context.Context, productID string) (*model.Cart, error) {
	userId, err := helper.Authenticated(ctx.Value("userId"))
	if err != nil {
		return nil, err
	}

	log.Println("Current user: ", userId)

	CartItems = append(CartItems, &model.CartItem{
		ID:          faker.UUIDDigit(),
		Title:       "Test Product",
		Description: "Lorem",
		Price:       17.5,
		Image:       "https://via.placeholder.com/150",
	})

	c := model.Cart{
		Items:      CartItems,
		ItemsCount: len(CartItems),
	}

	return &c, nil
}

func (r *mutationResolver) RemoveFromCart(ctx context.Context, productID string) (*model.Cart, error) {
	userId, err := helper.Authenticated(ctx.Value("userId"))
	if err != nil {
		return nil, err
	}

	log.Println("Current user: ", userId)

	CartItems = append([]*model.CartItem{}, CartItems[1:]...)

	c := model.Cart{
		Items:      CartItems,
		ItemsCount: len(CartItems),
	}

	return &c, nil
}

func (r *queryResolver) SayHello(ctx context.Context) (string, error) {
	return "Hello World", nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	productResp, err := r.ProductsClient.ListProducts(ctx, &product_grpc.ProductListRequest{})
	if err != nil {
		return nil, err
	}

	for _, product := range productResp.Products {
		products = append(products, &model.Product{
			ID:          product.Id,
			Title:       product.Title,
			Description: product.Description,
			Price:       float64(product.Price),
			Image:       product.Image,
		})
	}

	return products, nil
}

func (r *queryResolver) Cart(ctx context.Context) (*model.Cart, error) {
	userId, err := helper.Authenticated(ctx.Value("userId"))
	if err != nil {
		return nil, err
	}

	log.Println("Current user: ", userId)
	c := model.Cart{
		Items:      CartItems,
		ItemsCount: len(CartItems),
	}

	return &c, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
