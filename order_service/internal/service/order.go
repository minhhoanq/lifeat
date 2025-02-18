package service

import (
	"context"

	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/order_service/internal/dataaccess/database"
	pb "github.com/minhhoanq/lifeat/order_service/internal/generated/order_service"
)

type OrderService interface {
	CreateOrder(ctx context.Context, arg *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}

type orderService struct {
	l                 logger.Interface
	orderDataAccessor database.OrderDataAccessor
}

// catalogAccessor database.CatalogDataAccessor
func NewOrderService(l logger.Interface, orderDataAccessor database.OrderDataAccessor) OrderService {
	return &orderService{
		l:                 l,
		orderDataAccessor: orderDataAccessor,
	}
}

func (o *orderService) CreateOrder(ctx context.Context, arg *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &database.CreateOrderRequest{
		UserID:     arg.UserId,
		Address:    "",
		Status:     "PENDING",
		OrderItems: make([]database.CreateOrderItemRequest, 0, len(arg.CartItems)),
	}
	for _, item := range arg.CartItems {
		order.OrderItems = append(order.OrderItems, database.CreateOrderItemRequest{
			SkuID:    item.SkuId,
			Quantity: item.Quantity,
		})
	}

	result, err := o.orderDataAccessor.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:     result.Order.ID.String(),
			UserId: result.Order.UserID.String(),
			Status: "PENDING",
			Items:  make([]*pb.OrderItem, 0, len(arg.CartItems)),
		},
	}

	for _, item := range result.OrderItems {
		response.Order.Items = append(response.Order.Items, &pb.OrderItem{
			Id:       item.ID.String(),
			OrderId:  item.OrderID.String(),
			SkuId:    item.SkuID.String(),
			Quantity: item.Quantity,
		})
	}

	return nil, nil
}
