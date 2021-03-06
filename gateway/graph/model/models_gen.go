// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cart struct {
	Items []*CartItem `json:"items"`
}

type CartItem struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"productId"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

type FavouriteProduct struct {
	ID        string  `json:"id"`
	ProductID string  `json:"productID"`
	Status    int     `json:"status"`
	Title     *string `json:"title"`
	Image     *string `json:"image"`
}

type LoginResponse struct {
	ID       string `json:"id"`
	Avatar   string `json:"avatar"`
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}

type Order struct {
	PaymentAmount float64    `json:"paymentAmount"`
	CreatedAt     string     `json:"createdAt"`
	OrderItems    []*Product `json:"orderItems"`
}

type PaymentStartResponse struct {
	URL string `json:"url"`
}

type Product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}
