package grpc

import (
	"context"
	"fmt"

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
		l:              l,
	}, nil
}

func (h *Handler) CreateProduct(ctx context.Context, arg *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	h.l.Info("Create product in handler")
	fmt.Println("handle create product")
	product, err := h.catalogService.CreateProduct(ctx, arg)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// func (h *Handler) GetProductById(context.Context, *pb.GetProductByIdRequest) (*pb.GetProductByIdResponse, error) {
// 	return nil, nil
// }
