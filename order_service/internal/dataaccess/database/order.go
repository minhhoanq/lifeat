package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/minhhoanq/lifeat/common/logger"
	"gorm.io/gorm"
)

type OrderDataAccessor interface {
	CreateOrder(ctx context.Context, arg *CreateOrderRequest) (*CreateOrderResponse, error)
	createOrderItems(ctx context.Context, arg []CreateOrderItemRequest, tx *gorm.DB) ([]OrderItem, error)
}

type orderDataAccessor struct {
	database Database
	l        logger.Interface
}

func NewOrderDataAccessor(database Database, l logger.Interface) OrderDataAccessor {
	return &orderDataAccessor{
		database: database,
		l:        l,
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
		if err := tx.WithContext(ctx).Create(order).Error; err != nil {
			return err
		}

		orderItems, err := o.createOrderItems(ctx, arg.OrderItems, tx)

		response = &CreateOrderResponse{
			Order:      *order,
			OrderItems: orderItems,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, err
}

func (o *orderDataAccessor) createOrderItems(ctx context.Context, arg []CreateOrderItemRequest, tx *gorm.DB) ([]OrderItem, error) {
	orderItems := make([]OrderItem, 0, len(arg))

	for _, item := range arg {
		skuID, err := uuid.Parse(item.SkuID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse uuid with sku id")
		}
		orderID, err := uuid.Parse(item.OrderID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse uuid with order id")
		}
		orderItems = append(orderItems, OrderItem{
			SkuID:    skuID,
			OrderID:  orderID,
			Quantity: item.Quantity,
		})
	}

	if err := tx.WithContext(ctx).Create(&orderItems).Error; err != nil {
		return nil, err
	}

	return orderItems, nil
}
