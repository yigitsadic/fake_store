package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/yigitsadic/fake_store/auth/client/client"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/gateway/graph/generated"
	"github.com/yigitsadic/fake_store/gateway/graph/model"
	"github.com/yigitsadic/fake_store/gateway/helper"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

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

	product, err := r.ProductsClient.ProductDetail(ctx, &product_grpc.ProductDetailRequest{ProductId: productID})
	if err != nil {
		return nil, errors.New("product not found")
	}

	res, err := r.CartService.AddToCart(ctx, &cart_grpc.AddToCartRequest{
		UserId:      userId,
		ProductId:   product.Id,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
	})
	if err != nil {
		return nil, err
	}

	return &model.Cart{
		Items:      ConvertCartFromService(res.GetCartItems()),
		ItemsCount: int(res.GetItemCount()),
	}, nil
}

func (r *mutationResolver) RemoveFromCart(ctx context.Context, cartItemID string) (*model.Cart, error) {
	userId, err := helper.Authenticated(ctx.Value("userId"))
	if err != nil {
		return nil, err
	}

	res, err := r.CartService.RemoveFromCart(ctx, &cart_grpc.RemoveFromCartRequest{
		UserId:     userId,
		CartItemId: cartItemID,
	})
	if err != nil {
		return nil, err
	}

	return &model.Cart{
		Items:      ConvertCartFromService(res.GetCartItems()),
		ItemsCount: int(res.GetItemCount()),
	}, nil
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

	res, err := r.CartService.CartContent(ctx, &cart_grpc.CartContentRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	return &model.Cart{
		Items:      ConvertCartFromService(res.GetCartItems()),
		ItemsCount: int(res.GetItemCount()),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
