package service

import (
	"context"
	"fmt"

	pb "github.com/minhhoanq/lifeat/catalog_service/internal/generated/catalog_service"
	"github.com/minhhoanq/lifeat/common/logger"

	"gorm.io/gorm"
)

type CatalogService interface {
	CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.CreateProductResponse, error)
	GetProductById(context.Context, *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error)
}

type catalogService struct {
	db *gorm.DB
	l  logger.Interface
}

func NewCatalogService(db *gorm.DB, l logger.Interface) CatalogService {
	fmt.Println("config service")
	return &catalogService{
		db: db,
		l:  l,
	}
}

func (c *catalogService) CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return nil, nil
}
func (c *catalogService) GetProductById(context.Context, *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	return nil, nil
}
