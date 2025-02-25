package database

import (
	"time"

	"github.com/google/uuid"
)

// CreateProduct
type CreateProductParams struct {
	Name        string
	Description string
	Image       string
	CategoryID  int32
	BrandID     int32
	SKUs        []CreateSKUParams
}

type CreateSKUParams struct {
	Name       string
	Slug       string
	Price      CreatePriceParams
	Inventory  CreateInventoryParams
	Attributes []CreateAttributeParams
}

type CreatePriceParams struct {
	OriginalPrice int32
	EffectiveDate time.Time
}

type CreateInventoryParams struct {
	Stock int32
}

type CreateAttributeParams struct {
	AttributeID int32
	Value       string
}

// Response structures
type CreateProductResponse struct {
	ProductDetail ProductDetail
}

type ProductDetail struct {
	Product Product     `json:"product"`
	SKUs    []SKUDetail `json:"skus"`
}

type SKUDetail struct {
	SKU SKU `json:"sku"`
}

// ListProduct
type ListProductRequest struct {
	Page     int32
	PageSize int32
}

type ListProductResponse struct {
	Products []ProductDetail
}

type CreateCartRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type CreateCartResponse struct {
	CartID uuid.UUID `json:"cart_id"`
	UserID uuid.UUID `json:"user_id"`
}

type AddToCartItemRequest struct {
	CartID   uuid.UUID `json:"cart_id"`
	SkuID    uuid.UUID `json:"sku_id"`
	Quantity int32     `json:"quantity"`
}

type AddToCartItemResponse struct {
	Cart     Cart       `json:"cart"`
	CartItem []CartItem `json:"cart_item"`
}

type GetSKURequest struct {
	SkuID uuid.UUID `json:"sku_id"`
}

type GetSKUResponse struct {
	SKU SKU `json:"sku"`
}

type GetInventorySKURequest struct {
	SkuID uuid.UUID `json:"sku_id"`
}

type GetInventorySKUResponse struct {
	Inventory Inventory `json:"inventory"`
}

type UpdateInventorySKURequest struct {
	SkuID    uuid.UUID `json:"sku_id"`
	Quantity int32     `json:"quantitty"`
}

type UpdateInventorySKUResponse struct {
	Inventory Inventory `json:"inventory"`
}
