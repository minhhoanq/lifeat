package catalogservice

import (
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/order_service/config"
	"github.com/minhhoanq/lifeat/order_service/internal/generated/catalog_service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(cfg config.Config, l logger.Interface) (catalog_service.CatalogServiceClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.NewClient(cfg.GRPCCatalogAddress, opts...)
	if err != nil {
		l.Error("failed to connect catalog grpc client", zap.Error(err))
		return nil, err
	}

	client := catalog_service.NewCatalogServiceClient(conn)

	return client, nil
}
