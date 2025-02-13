package database

import (
	"context"

	"github.com/minhhoanq/lifeat/common/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Models for database entities

// Request/Response structures

// Interface definition
type CatalogDataAccessor interface {
	CreateProduct(ctx context.Context, arg *CreateProductParams) (*CreateProductResponse, error)
	ListProdudct(ctx context.Context, arg *ListProductRequest) (*ListProductResponse, error)
}

type catalogDataAccessor struct {
	database Database
	l        logger.Interface
}

func NewCatalogDataAccessor(database Database, l logger.Interface) CatalogDataAccessor {
	l.Info("initializing catalog data accessor")
	return &catalogDataAccessor{
		database: database,
		l:        l,
	}
}

func (c *catalogDataAccessor) CreateProduct(ctx context.Context, arg *CreateProductParams) (*CreateProductResponse, error) {
	c.l.Info("starting product creation transaction", zap.String("product_name", arg.Name))

	var response *CreateProductResponse

	// Start transaction
	err := c.database.Transaction(func(tx *gorm.DB) error {
		// Create product
		product := Product{
			Name:        arg.Name,
			Description: arg.Description,
			Image:       arg.Image,
			CategoryID:  arg.CategoryID,
			BrandID:     arg.BrandID,
		}

		if err := tx.Create(&product).Error; err != nil {
			c.l.Error("failed to create product",
				zap.String("product_name", arg.Name),
				zap.Error(err))
			return err
		}

		var skuDetails []SKUDetail
		// Create SKUs and related entities
		for _, skuParam := range arg.SKUs {
			skuDetail, err := c.createSKUWithDetails(tx, &product, &skuParam)
			if err != nil {
				return err
			}
			skuDetails = append(skuDetails, *skuDetail)
		}

		response = &CreateProductResponse{
			ProductDetail: ProductDetail{
				Product: product,
				SKUs:    skuDetails,
			},
		}

		return nil
	})

	if err != nil {
		c.l.Error("transaction failed", zap.Error(err))
		return nil, err
	}

	c.l.Info("product creation completed successfully",
		zap.String("product_id", response.ProductDetail.Product.ID.String()))

	return response, nil
}

func (c *catalogDataAccessor) createSKUWithDetails(
	tx *gorm.DB,
	product *Product,
	skuParam *CreateSKUParams,
) (*SKUDetail, error) {
	// Create SKU
	sku := SKU{
		ProductID: product.ID,
		Name:      skuParam.Name,
		Slug:      skuParam.Slug,
	}

	if err := tx.Create(&sku).Error; err != nil {
		c.l.Error("failed to create SKU",
			zap.String("sku_name", skuParam.Name),
			zap.Error(err))
		return nil, err
	}

	// Create price
	price := Price{
		SkuID:         sku.ID,
		OriginalPrice: skuParam.Price.OriginalPrice,
		EffectiveDate: skuParam.Price.EffectiveDate,
		Active:        true,
	}

	if err := tx.Create(&price).Error; err != nil {
		c.l.Error("failed to create price",
			zap.String("sku_id", sku.ID.String()),
			zap.Error(err))
		return nil, err
	}

	// Create inventory
	inventory := Inventory{
		SkuID: sku.ID,
		Stock: skuParam.Inventory.Stock,
	}

	if err := tx.Create(&inventory).Error; err != nil {
		c.l.Error("failed to create inventory",
			zap.String("sku_id", sku.ID.String()),
			zap.Error(err))
		return nil, err
	}

	// Create SKU attributes
	var skuAttributes []SKUAttribute
	for _, attrParam := range skuParam.Attributes {
		skuAttr := SKUAttribute{
			SkuID:       sku.ID,
			AttributeID: attrParam.AttributeID,
			Value:       attrParam.Value,
		}

		if err := tx.Create(&skuAttr).Error; err != nil {
			c.l.Error("failed to create SKU attribute",
				zap.String("sku_id", sku.ID.String()),
				zap.Int32("attribute_id", attrParam.AttributeID),
				zap.Error(err))
			return nil, err
		}

		skuAttributes = append(skuAttributes, skuAttr)
	}

	return &SKUDetail{
		SKU:           sku,
		Price:         price,
		Inventory:     inventory,
		SKUAttributes: skuAttributes,
	}, nil
}

func (c *catalogDataAccessor) ListProdudct(ctx context.Context, arg *ListProductRequest) (*ListProductResponse, error) {
	return nil, nil
}
