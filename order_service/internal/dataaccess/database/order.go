package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/order_service/internal/generated/catalog_service"
	"gorm.io/gorm"
)

type OrderDataAccessor interface {
	CreateOrder(ctx context.Context, arg *CreateOrderRequest) (*CreateOrderResponse, error)
	createOrderItems(ctx context.Context, order_id uuid.UUID, arg []CreateOrderItemRequest, tx *gorm.DB) ([]OrderItem, error)
}

type orderDataAccessor struct {
	database             Database
	l                    logger.Interface
	catalogServiceClient catalog_service.CatalogServiceClient
}

func NewOrderDataAccessor(database Database, l logger.Interface, catalogServiceClient catalog_service.CatalogServiceClient) OrderDataAccessor {
	return &orderDataAccessor{
		database:             database,
		l:                    l,
		catalogServiceClient: catalogServiceClient,
	}
}

func (o *orderDataAccessor) CreateOrder(ctx context.Context, arg *CreateOrderRequest) (*CreateOrderResponse, error) {
	var response *CreateOrderResponse

	err := o.database.Transaction(func(tx *gorm.DB) error {
		userID, err := uuid.Parse(arg.UserID)
		if err != nil {
			return err
		}

		order := &Order{
			UserID:  userID,
			Address: arg.Address,
			Status:  arg.Status,
		}
		if err := tx.WithContext(ctx).Create(&order).Error; err != nil {
			return err
		}

		orderItems, err := o.createOrderItems(ctx, order.ID, arg.OrderItems, tx)

		response = &CreateOrderResponse{
			Order: Order{
				ID:            order.ID,
				UserID:        order.UserID,
				Address:       order.Address,
				Status:        order.Status,
				PaymentMethod: order.PaymentMethod,
				CreatedAt:     order.CreatedAt,
				UpdatedAt:     order.UpdatedAt,
			},
			OrderItems: orderItems,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, err
}

func (o *orderDataAccessor) createOrderItems(ctx context.Context, order_id uuid.UUID, arg []CreateOrderItemRequest, tx *gorm.DB) ([]OrderItem, error) {
	orderItems := make([]OrderItem, 0, len(arg))

	for _, item := range arg {
		skuID, err := uuid.Parse(item.SkuID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse uuid with sku id")
		}

		orderItems = append(orderItems, OrderItem{
			SkuID:    skuID,
			OrderID:  order_id,
			Quantity: item.Quantity,
		})
		o.catalogServiceClient.UpdateInventorySKU(ctx, &catalog_service.UpdateInventorySKURequest{
			SkuId:    item.SkuID,
			Quantity: item.Quantity,
		})
	}

	if err := tx.WithContext(ctx).Create(&orderItems).Error; err != nil {
		return nil, err
	}

	return orderItems, nil
}
