package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/yigitsadic/fake_store/auth/client/client"
	"github.com/yigitsadic/fake_store/gateway/graph/generated"
	"github.com/yigitsadic/fake_store/gateway/graph/model"
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

func (r *queryResolver) SayHello(ctx context.Context) (string, error) {
	return "Hello World", nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products := []*model.Product{
		{
			ID:          "12e",
			Title:       "Camera",
			Description: "Basic camera",
			Price:       499.50,
			Image:       "camera.png",
		},
		{
			ID:          "24e",
			Title:       "Game console",
			Description: "Game console. Video games.",
			Price:       300.75,
			Image:       "game-console.png",
		},
		{
			ID:          "43ert5",
			Title:       "Classical Novel",
			Description: "Classical novel that we all like",
			Price:       27,
			Image:       "book-cover-1.png",
		},
	}

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
