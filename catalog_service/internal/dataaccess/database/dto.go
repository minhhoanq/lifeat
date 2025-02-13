package database

import "time"

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
	Product Product
	SKUs    []SKUDetail
}

type SKUDetail struct {
	SKU           SKU
	Price         Price
	Inventory     Inventory
	SKUAttributes []SKUAttribute
}
