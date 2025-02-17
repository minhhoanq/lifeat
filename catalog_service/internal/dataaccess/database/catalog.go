package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/minhhoanq/lifeat/common/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Models for database entities

// Request/Response structures

// Interface definition
type CatalogDataAccessor interface {
	CreateProduct(ctx context.Context, arg *CreateProductParams) (*CreateProductResponse, error)
	ListProducts(ctx context.Context, arg *ListProductRequest) (*ListProductResponse, error)
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
		SKU: sku,
	}, nil
}

func (c *catalogDataAccessor) ListProducts(ctx context.Context, arg *ListProductRequest) (*ListProductResponse, error) {
	var products []ProductDetail
	var rawResult []byte

	offset := (arg.Page - 1) * arg.PageSize

	fmt.Println("page", arg.Page)
	fmt.Println("page_size", arg.PageSize)
	fmt.Println("offset", offset)

	query := `
		WITH limited_products AS (
			SELECT * FROM products
			ORDER BY created_at DESC
			LIMIT ?
			OFFSET ?
		)
		SELECT json_agg(
			json_build_object(
				'product', json_build_object(
					'id', p.id,
					'name', p.name,
					'description', p.description,
					'image', p.image,
					'category_id', p.category_id,
					'brand_id', p.brand_id,
					'created_at', to_char(p.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
					'updated_at', to_char(p.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
				),
				'skus', (
					SELECT json_agg(
						json_build_object(
							'sku', json_build_object(
								'id', s.id,
								'product_id', s.product_id,
								'name', s.name,
								'slug', s.slug,
								'price', json_build_object(
									'id', pr.id,
									'sku_id', pr.sku_id,
									'original_price', pr.original_price,
									'effective_date', to_char(pr.effective_date, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
									'active', pr.active
								),
								'inventory', json_build_object(
									'id', i.id,
									'sku_id', i.sku_id,
									'stock', i.stock,
									'reservations', i.reservations
								),
								'attributes', (
									SELECT json_agg(
										json_build_object(
											'attribute_id', sa.attribute_id,
											'value', sa.value
										)
									)
									FROM sku_attributes sa
									WHERE sa.sku_id = s.id
								)
							)
						)
					)
					FROM skus s
					INNER JOIN prices pr ON s.id = pr.sku_id
					INNER JOIN inventories i ON s.id = i.sku_id
					WHERE s.product_id = p.id
				)
			)
		) AS result
		FROM limited_products p;`

	if err := c.database.Raw(query, arg.PageSize, offset).Row().Scan(&rawResult); err != nil {
		c.l.Error("failed to list products", zap.Error(err))
		return &ListProductResponse{
			Products: []ProductDetail{},
		}, err
	}

	if rawResult == nil {
		return &ListProductResponse{
			Products: []ProductDetail{},
		}, nil
	}

	if err := json.Unmarshal(rawResult, &products); err != nil {
		c.l.Error("failed to unmarshal products", zap.Error(err))
		return nil, err
	}

	return &ListProductResponse{
		Products: products,
	}, nil
}

// Helper method to get a single product with all its details
// func (r *ProductRepository) GetProductByID(ctx context.Context, productID uuid.UUID) (*ProductDetail, error) {
// 	var product Product
// 	if err := c.database.First(&product, "id = ?", productID).Error; err != nil {
// 		return nil, err
// 	}

// 	productDetail := ProductDetail{
// 		Product: product,
// 	}

// 	var skus []SKU
// 	if err := c.database.Where("product_id = ?", productID).Find(&skus).Error; err != nil {
// 		return nil, err
// 	}

// 	for _, sku := range skus {
// 		skuDetail := SKUDetail{
// 			SKU: sku,
// 		}

// 		// Get active price
// 		var price Price
// 		if err := c.database.Where("sku_id = ? AND active = true", sku.ID).
// 			Order("effective_date DESC").
// 			First(&price).Error; err != nil && err != gorm.ErrRecordNotFound {
// 			return nil, err
// 		}
// 		skuDetail.Price = price

// 		// Get inventory
// 		var inventory Inventory
// 		if err := c.database.Where("sku_id = ?", sku.ID).First(&inventory).Error; err != nil && err != gorm.ErrRecordNotFound {
// 			return nil, err
// 		}
// 		skuDetail.Inventory = inventory

// 		// Get attributes
// 		var attributes []SKUAttribute
// 		if err := c.database.Where("sku_id = ?", sku.ID).Find(&attributes).Error; err != nil {
// 			return nil, err
// 		}
// 		skuDetail.SKUAttributes = attributes

// 		productDetail.SKUs = append(productDetail.SKUs, skuDetail)
// 	}

// 	return &productDetail, nil
// }
