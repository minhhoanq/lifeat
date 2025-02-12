package initial

import (
	"github.com/minhhoanq/lifeat/catalog_service/configs"
	"github.com/minhhoanq/lifeat/catalog_service/internal/dataaccess/database"
	"github.com/minhhoanq/lifeat/catalog_service/internal/handler/grpc"
	"github.com/minhhoanq/lifeat/catalog_service/internal/service"
	"github.com/minhhoanq/lifeat/common/logger"
)

func InitialServer(cfg configs.Config, l logger.Interface) (grpc.Server, error) {

	database, err := database.New(cfg, l)
	if err != nil {
		return nil, err
	}

	catalogService := service.NewCatalogService(database.DB, l)
	handler, err := grpc.NewHandler(catalogService, l)
	if err != nil {
		return nil, err
	}
	server := grpc.NewGRPCServer(cfg, handler, l)

	return server, nil
}
