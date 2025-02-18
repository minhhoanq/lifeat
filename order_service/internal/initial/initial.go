package initial

import (
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/order_service/config"
	"github.com/minhhoanq/lifeat/order_service/internal/dataaccess/database"
	"github.com/minhhoanq/lifeat/order_service/internal/handler/grpc"
	"github.com/minhhoanq/lifeat/order_service/internal/service"
)

func InitialServer(cfg config.Config, l logger.Interface) (grpc.Server, error) {
	db, err := database.New(cfg, l)
	if err != nil {
		return nil, err
	}

	orderDataAccessor := database.NewOrderDataAccessor(db, l)
	// Initialize the service
	orderService := service.NewOrderService(l, orderDataAccessor)
	// Initialize the handler
	handler, err := grpc.NewHandler(l, orderService)
	// Initialize the server
	grpcServer := grpc.NewGRPCServer(cfg, handler, l)

	return grpcServer, nil
}
