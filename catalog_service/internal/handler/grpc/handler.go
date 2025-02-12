package grpc

import (
	"context"

	pb "github.com/minhhoanq/lifeat/catalog_service/internal/generated/catalog_service"
	"github.com/minhhoanq/lifeat/catalog_service/internal/service"
	"github.com/minhhoanq/lifeat/common/logger"
)

type Handler struct {
	pb.UnimplementedCatalogServiceServer
	catalogService service.CatalogService
	l              logger.Interface
}

func NewHandler(catalogService service.CatalogService, l logger.Interface) (pb.CatalogServiceServer, error) {
	return &Handler{
		catalogService: catalogService,
	}, nil
}

func (h *Handler) CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return nil, nil
}
func (h *Handler) GetProductById(context.Context, *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
	return nil, nil
}
