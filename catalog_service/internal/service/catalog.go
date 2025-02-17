package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/minhhoanq/lifeat/catalog_service/internal/dataaccess/database"
	pb "github.com/minhhoanq/lifeat/catalog_service/internal/generated/catalog_service"
	"github.com/minhhoanq/lifeat/common/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gorm"
)

type CatalogService interface {
	CreateProduct(ctx context.Context, arg *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	ListProduct(ctx context.Context, arg *pb.ListProductRequest) (*pb.ListProductResponse, error)
	CreateCart(ctx context.Context, arg *pb.CreateCartRequest) (*pb.CreateCartResponse, error)
	AddToCartItem(ctx context.Context, arg *pb.AddToCartItemRequest) (*pb.AddToCartItemResponse, error)
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

func convertProductResponse(dbResponse *database.ProductDetail) *pb.ProductWithSKUs {
	// Transform database response to protobuf response
	response := &pb.ProductWithSKUs{
		Product: &pb.Product{
			Id:          dbResponse.Product.ID.String(),
			Name:        dbResponse.Product.Name,
			Description: dbResponse.Product.Description,
			Image:       dbResponse.Product.Image,
			CategoryId:  dbResponse.Product.CategoryID,
			BrandId:     dbResponse.Product.BrandID,
			CreatedAt:   timestamppb.New(dbResponse.Product.CreatedAt),
			UpdatedAt:   timestamppb.New(dbResponse.Product.UpdatedAt),
		},
		Skus: make([]*pb.SKU, 0, len(dbResponse.SKUs)),
	}

	// Transform SKUs in response
	for _, skuDetail := range dbResponse.SKUs {
		sku := &pb.SKU{
			Id:        skuDetail.SKU.ID.String(),
			ProductId: skuDetail.SKU.ProductID.String(),
			Name:      skuDetail.SKU.Name,
			Slug:      skuDetail.SKU.Slug,
			CreatedAt: timestamppb.New(skuDetail.SKU.CreatedAt),
			UpdatedAt: timestamppb.New(skuDetail.SKU.UpdatedAt),
			CurrentPrice: &pb.Price{
				Id:            skuDetail.SKU.Price.ID,
				SkuId:         skuDetail.SKU.Price.SkuID.String(),
				OriginalPrice: skuDetail.SKU.Price.OriginalPrice,
				EffectiveDate: timestamppb.New(skuDetail.SKU.Price.EffectiveDate),
				Active:        skuDetail.SKU.Price.Active,
			},
			Inventory: &pb.Inventory{
				Id:           skuDetail.SKU.Inventory.ID,
				SkuId:        skuDetail.SKU.Inventory.SkuID.String(),
				Stock:        skuDetail.SKU.Inventory.Stock,
				Reservations: string(skuDetail.SKU.Inventory.Reservations),
			},
			Attributes: make([]*pb.AttributeValue, 0, len(skuDetail.SKU.SKUAttributes)),
		}

		// Transform SKU attributes
		for _, attr := range skuDetail.SKU.SKUAttributes {
			sku.Attributes = append(sku.Attributes, &pb.AttributeValue{
				AttributeId: attr.AttributeID,
				Value:       attr.Value,
			})
		}

		response.Skus = append(response.Skus, sku)
	}

	return response
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
				EffectiveDate: time.Now(),
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

	response := convertProductResponse(&dbResponse.ProductDetail)

	c.l.Info("product created successfully",
		zap.String("product_id", response.Product.Id))

	return &pb.CreateProductResponse{
		Product: response,
	}, nil
}

func (c *catalogService) ListProduct(ctx context.Context, arg *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	products, err := c.catalogAccessor.ListProducts(ctx, &database.ListProductRequest{
		Page:     arg.GetPage(),
		PageSize: arg.GetPageSize(),
	})
	if err != nil {
		return nil, err
	}
	c.l.Info("len", zap.Int("len", len(products.Products)))
	response := make([]*pb.ProductWithSKUs, 0, len(products.Products))

	for _, product := range products.Products {
		response = append(response, convertProductResponse(&product))
	}

	return &pb.ListProductResponse{
		Products: response,
	}, nil
}

func (c *catalogService) CreateCart(ctx context.Context, arg *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	uuidUserID, err := uuid.Parse(arg.UserId)
	if err != nil {
		c.l.Error("failed to parse user id", zap.Error(err))
		return nil, err
	}

	cart, err := c.catalogAccessor.CreateCart(ctx, &database.CreateCartRequest{
		UserID: uuidUserID,
	})
	if err != nil {
		c.l.Error("failed to create cart", zap.Error(err))
		return nil, err
	}

	return &pb.CreateCartResponse{
		CartId: cart.CartID.String(),
		UserId: cart.UserID.String(),
	}, nil
}

func (c *catalogService) AddToCartItem(ctx context.Context, arg *pb.AddToCartItemRequest) (*pb.AddToCartItemResponse, error) {
	uuidCartID, err := uuid.Parse(arg.Item.CartId)
	if err != nil {
		c.l.Error("failed to parse cart id", zap.Error(err))
		return nil, err
	}

	uuidSkuID, err := uuid.Parse(arg.Item.SkuId)
	if err != nil {
		c.l.Error("failed to parse sku id", zap.Error(err))
		return nil, err
	}

	cartItems, err := c.catalogAccessor.AddToCartItem(ctx, &database.AddToCartItemRequest{
		CartID:   uuidCartID,
		SkuID:    uuidSkuID,
		Quantity: arg.Item.Quantity,
	})

	if err != nil {
		c.l.Error("failed to add item to cart", zap.Error(err))
		return nil, err
	}

	response := &pb.AddToCartItemResponse{
		CartId: cartItems.Cart.ID.String(),
		Items:  make([]*pb.CartItem, 0, len(cartItems.CartItem)),
	}

	for _, item := range cartItems.CartItem {
		response.Items = append(response.Items, &pb.CartItem{
			Id:       item.ID.String(),
			CartId:   item.CartID.String(),
			SkuId:    item.SkuID.String(),
			Quantity: item.Quantity,
		})
	}

	return response, nil
}
