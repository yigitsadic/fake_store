package event_handlers

import "encoding/json"

type populateProductRequestMessage struct {
	ProductID  string `json:"product_id"`
	CartItemID string `json:"cart_item_id"`
}

type CartItemProductMessage struct {
	ProductID   string  `json:"product_id"`
	CartItemID  string  `json:"cart_item_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float32 `json:"price"`
}

func unmarshalProductDataRequest(payload string) (populateProductRequestMessage, error) {
	var res populateProductRequestMessage

	err := json.Unmarshal([]byte(payload), &res)
	if err != nil {
		return populateProductRequestMessage{}, err
	}

	return res, nil
}
