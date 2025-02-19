package database

type CreateOrderRequest struct {
	UserID     string                   `json:"user_id"`
	Address    string                   `json:"address"`
	Status     string                   `json:"status"`
	OrderItems []CreateOrderItemRequest `json:"create_order_item_request"`
}

type CreateOrderResponse struct {
	Order      Order       `json:"order"`
	OrderItems []OrderItem `json:"order_items"`
}

type CreateOrderItemRequest struct {
	SkuID    string `json:"sku_id"`
	Quantity int32  `json:"quantity"`
}
