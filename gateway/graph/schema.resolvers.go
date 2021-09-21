package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/yigitsadic/fake_store/auth/auth_grpc/auth_grpc"
	"github.com/yigitsadic/fake_store/cart/cart_grpc/cart_grpc"
	"github.com/yigitsadic/fake_store/favourites/favourites_grpc/favourites_grpc"
	"github.com/yigitsadic/fake_store/gateway/graph/generated"
	"github.com/yigitsadic/fake_store/gateway/graph/model"
	"github.com/yigitsadic/fake_store/gateway/middlewares"
	"github.com/yigitsadic/fake_store/orders/orders_grpc/orders_grpc"
	"github.com/yigitsadic/fake_store/products/product_grpc/product_grpc"
)

func (r *mutationResolver) Login(ctx context.Context) (*model.LoginResponse, error) {
	result, err := r.AuthService.LoginUser(ctx, &auth_grpc.AuthRequest{})
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

func (r *mutationResolver) AddToCart(ctx context.Context, productID string) (bool, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return false, err
	}

	_, err = r.CartService.AddToCart(ctx, &cart_grpc.AddToCartRequest{
		UserId:    userID,
		ProductId: productID,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) RemoveFromCart(ctx context.Context, cartItemID string) (bool, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return false, err
	}

	_, err = r.CartService.RemoveFromCart(ctx, &cart_grpc.RemoveFromCartRequest{
		UserId:     userID,
		CartItemId: cartItemID,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) StartPayment(ctx context.Context) (*model.PaymentStartResponse, error) {
	var err error

	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return nil, err
	}

	var paymentURL string

	// Get cart
	cart, err := r.CartService.CartContent(ctx, &cart_grpc.CartContentRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	// Create and get reference id and payment total
	var cartItems []*orders_grpc.CartItem

	for _, item := range cart.GetCartItems() {
		cartItems = append(cartItems, &orders_grpc.CartItem{
			Id:          item.GetId(),
			ProductId:   item.GetProductId(),
			Title:       item.GetTitle(),
			Description: item.GetDescription(),
			Price:       item.GetPrice(),
			Image:       item.GetImage(),
		})
	}

	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	paymentReq := &orders_grpc.StartOrderRequest{
		UserId:    userID,
		CartItems: cartItems,
	}

	res, err := r.OrdersService.StartOrder(ctx, paymentReq)
	if err != nil {
		return nil, err
	}

	operation := func() error {
		paymentURL, err = createPaymentIntent(ctx, r.PaymentProviderURL, paymentIntentRequest{
			Amount:      float64(res.GetPaymentAmount()),
			ReferenceID: res.GetId(),
			HookURL:     r.HookURL,
			SuccessURL:  r.SuccessURL,
			FailureURL:  r.FailureURL,
		})

		return err
	}

	err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
	if err != nil {
		return nil, err
	}

	return &model.PaymentStartResponse{URL: paymentURL}, nil
}

func (r *mutationResolver) AddToFavourites(ctx context.Context, productID string) (bool, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return false, err
	}

	res, err := r.FavouritesService.MarkFavourite(ctx, &favourites_grpc.FavouritesRequest{
		ProductID: productID,
		UserID:    userID,
	})
	if err != nil {
		return false, err
	}

	return res.GetSuccess(), nil
}

func (r *mutationResolver) RemoveFromFavourites(ctx context.Context, productID string) (bool, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return false, err
	}

	res, err := r.FavouritesService.UnMarkFavourite(ctx, &favourites_grpc.FavouritesRequest{
		ProductID: productID,
		UserID:    userID,
	})
	if err != nil {
		return false, err
	}

	return res.GetSuccess(), nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	productResp, err := r.ProductsService.ListProducts(ctx, &product_grpc.ProductListRequest{})
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

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	p, err := r.ProductsService.ProductDetail(ctx, &product_grpc.ProductDetailRequest{ProductId: id})
	if err != nil {
		return nil, err
	}

	product := &model.Product{
		ID:          p.GetId(),
		Title:       p.GetTitle(),
		Description: p.GetDescription(),
		Price:       float64(p.GetPrice()),
		Image:       p.GetImage(),
	}

	// if user logged in, we'll continue to add favourite
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		res, err := r.FavouritesService.ProductInFavourite(ctx, &favourites_grpc.FavouritesRequest{
			ProductID: id,
			UserID:    userID,
		})
		if err == nil {
			product.InFavourites = res.GetInFavourites()
		}
	}

	return product, nil
}

func (r *queryResolver) FavouriteProducts(ctx context.Context) ([]*model.FavouriteProduct, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return nil, err
	}

	var products []*model.FavouriteProduct

	result, err := r.FavouritesService.ListFavourites(ctx, &favourites_grpc.ListFavouritesRequest{UserID: userID})
	if err != nil {
		return nil, err
	}

	for _, item := range result.GetProducts() {
		var title, image string

		title = item.GetTitle()
		image = item.GetImage()

		products = append(products, &model.FavouriteProduct{
			ID:        item.GetId(),
			ProductID: item.GetProductID(),
			Title:     &title,
			Image:     &image,
		})
	}

	return products, nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.OrdersService.ListOrders(ctx, &orders_grpc.OrderListRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	var orderList []*model.Order

	for _, order := range res.GetOrders() {
		var products []*model.Product

		for _, product := range order.GetProducts() {
			products = append(products, &model.Product{
				ID:          product.GetId(),
				Title:       product.GetTitle(),
				Description: product.GetDescription(),
				Price:       float64(product.GetPrice()),
				Image:       product.GetImage(),
			})
		}

		orderList = append(orderList, &model.Order{
			PaymentAmount: float64(order.GetPaymentAmount()),
			CreatedAt:     order.GetCreatedAt(),
			OrderItems:    products,
		})
	}

	return orderList, nil
}

func (r *queryResolver) Cart(ctx context.Context) (*model.Cart, error) {
	userID, err := middlewares.Authenticated(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.CartService.CartContent(ctx, &cart_grpc.CartContentRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	return &model.Cart{
		Items: convertCartFromService(res.GetCartItems()),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
