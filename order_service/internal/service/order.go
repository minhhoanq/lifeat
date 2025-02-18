package service

import (
	"github.com/minhhoanq/lifeat/common/logger"

	"gorm.io/gorm"
)

type OrderService interface {
}

type orderService struct {
	db *gorm.DB
	l  logger.Interface
	// catalogAccessor   database.CatalogDataAccessor
}

// catalogAccessor database.CatalogDataAccessor
func NewOrderService(db *gorm.DB, l logger.Interface) OrderService {
	return &orderService{
		db: db,
		l:  l,
	}
}
