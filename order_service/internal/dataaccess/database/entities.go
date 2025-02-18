package database

import "github.com/google/uuid"

type Order struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID  uuid.UUID `json:"user_id"`
	Address string    `json:"address"`
	Status  string    `json:"status"`
}

type OrderItem struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrderID  uuid.UUID `json:"order_id"`
	SkuID    uuid.UUID `json:"sku_id"`
	Quantity int32     `json:"quantity"`
}
