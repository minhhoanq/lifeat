package service

import (
	"context"
	"time"

	"github.com/minhhoanq/lifeat/catalog_service/internal/dataaccess/database"
	pb "github.com/minhhoanq/lifeat/catalog_service/internal/generated/catalog_service"
	"github.com/minhhoanq/lifeat/common/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gorm"
)

type CatalogService interface {
	CreateProduct(ctx context.Context, arg *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	// GetProductById(context.Context, *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error)
}

type catalogService struct {
	db              *gorm.DB
	l               logger.Interface
	catalogAccessor database.CatalogDataAccessor
}

func NewCatalogService(db *gorm.DB, l logger.Interface, catalogAccessor database.CatalogDataAccessor) CatalogService {
	return &catalogService{
		db:              db,
		l:               l,
		catalogAccessor: catalogAccessor,
	}
}

func (c *catalogService) CreateProduct(ctx context.Context, arg *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	c.l.Info("Create product in service")
	// Transform protobuf request to database parameters
	dbParams := &database.CreateProductParams{
		Name:        arg.Name,
		Description: arg.Description,
		Image:       arg.Image,
		CategoryID:  arg.CategoryId,
		BrandID:     arg.BrandId,
		SKUs:        make([]database.CreateSKUParams, 0, len(arg.Skus)),
	}

	// Transform SKU parameters
	for _, skuProto := range arg.Skus {
		skuParams := database.CreateSKUParams{
			Name: skuProto.Name,
			Slug: skuProto.Slug,
			Price: database.CreatePriceParams{
				OriginalPrice: skuProto.OriginalPrice,
				EffectiveDate: time.Now(), // You might want to get this from the request
			},
			Inventory: database.CreateInventoryParams{
				Stock: skuProto.InitialStock,
			},
			Attributes: make([]database.CreateAttributeParams, 0, len(skuProto.Attributes)),
		}

		// Transform attributes
		for _, attrProto := range skuProto.Attributes {
			skuParams.Attributes = append(skuParams.Attributes, database.CreateAttributeParams{
				AttributeID: attrProto.AttributeId,
				Value:       attrProto.Value,
			})
		}

		dbParams.SKUs = append(dbParams.SKUs, skuParams)
	}

	// Call database accessor
	dbResponse, err := c.catalogAccessor.CreateProduct(ctx, dbParams)
	if err != nil {
		c.l.Error("failed to create product in database",
			zap.String("product_name", arg.Name),
			zap.Error(err))
		return nil, err
	}

	// Transform database response to protobuf response
	response := &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          dbResponse.ProductDetail.Product.ID.String(),
			Name:        dbResponse.ProductDetail.Product.Name,
			Description: dbResponse.ProductDetail.Product.Description,
			Image:       dbResponse.ProductDetail.Product.Image,
			CategoryId:  dbResponse.ProductDetail.Product.CategoryID,
			BrandId:     dbResponse.ProductDetail.Product.BrandID,
			CreatedAt:   timestamppb.New(dbResponse.ProductDetail.Product.CreatedAt),
			UpdatedAt:   timestamppb.New(dbResponse.ProductDetail.Product.UpdatedAt),
		},
		Skus: make([]*pb.SKU, 0, len(dbResponse.ProductDetail.SKUs)),
	}

	// Transform SKUs in response
	for _, skuDetail := range dbResponse.ProductDetail.SKUs {
		sku := &pb.SKU{
			Id:        skuDetail.SKU.ID.String(),
			ProductId: skuDetail.SKU.ProductID.String(),
			Name:      skuDetail.SKU.Name,
			Slug:      skuDetail.SKU.Slug,
			CreatedAt: timestamppb.New(skuDetail.SKU.CreatedAt),
			UpdatedAt: timestamppb.New(skuDetail.SKU.UpdatedAt),
			CurrentPrice: &pb.Price{
				Id:            skuDetail.Price.ID,
				SkuId:         skuDetail.Price.SkuID.String(),
				OriginalPrice: skuDetail.Price.OriginalPrice,
				EffectiveDate: timestamppb.New(skuDetail.Price.EffectiveDate),
				Active:        skuDetail.Price.Active,
			},
			Inventory: &pb.Inventory{
				Id:    skuDetail.Inventory.ID,
				SkuId: skuDetail.Inventory.SkuID.String(),
				Stock: skuDetail.Inventory.Stock,
				// Reservations: skuDetail.Inventory.Reservations,
			},
			Attributes: make([]*pb.AttributeValue, 0, len(skuDetail.SKUAttributes)),
		}

		// Transform SKU attributes
		for _, attr := range skuDetail.SKUAttributes {
			sku.Attributes = append(sku.Attributes, &pb.AttributeValue{
				AttributeId: attr.AttributeID,
				Value:       attr.Value,
			})
		}

		response.Skus = append(response.Skus, sku)
	}

	c.l.Info("product created successfully",
		zap.String("product_id", response.Product.Id))

	return response, nil
}

// func (c *catalogService) GetProductById(context.Context, *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
// 	return nil, nil
// }
