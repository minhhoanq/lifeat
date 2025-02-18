package grpc

import (
	"github.com/minhhoanq/lifeat/common/logger"
	pb "github.com/minhhoanq/lifeat/order_service/internal/generated/order_service"
	"github.com/minhhoanq/lifeat/order_service/internal/service"
)

type Handler struct {
	pb.UnimplementedOrderServiceServer
	l            logger.Interface
	orderService service.OrderService
}

func NewHandler(l logger.Interface, orderService service.OrderService) (pb.OrderServiceServer, error) {
	return &Handler{
		l:            l,
		orderService: orderService,
	}, nil
}
