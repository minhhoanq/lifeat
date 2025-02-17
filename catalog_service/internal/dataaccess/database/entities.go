package database

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CategoryID  int32     `json:"category_id"`
	BrandID     int32     `json:"brand_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SKU struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ProductID     uuid.UUID      `json:"product_id"`
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	Price         Price          `gorm:"foreignKey:SkuID"`
	Inventory     Inventory      `gorm:"foreignKey:SkuID"`
	SKUAttributes []SKUAttribute `gorm:"foreignKey:SkuID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Price struct {
	ID            int32     `gorm:"primaryKey;autoIncrement"`
	SkuID         uuid.UUID `json:"sku_id"`
	OriginalPrice int32     `json:"original_price"`
	EffectiveDate time.Time `json:"effective_date"`
	Active        bool      `json:"active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Inventory struct {
	ID           int32           `gorm:"primaryKey;autoIncrement"`
	SkuID        uuid.UUID       `json:"sku_id"`
	Stock        int32           `json:"stock"`
	Reservations json.RawMessage `gorm:"type:json"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type SKUAttribute struct {
	ID          int32     `gorm:"primaryKey;autoIncrement"`
	SkuID       uuid.UUID `json:"attr_sku_id"`
	AttributeID int32     `json:"attribute_id"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Cart struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID uuid.UUID `json:"user_id"`
}

type CartItem struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CartID   uuid.UUID `json:"cart_id"`
	SkuID    uuid.UUID `json:"sku_id"`
	Quantity int32     `json:"quantity"`
}
