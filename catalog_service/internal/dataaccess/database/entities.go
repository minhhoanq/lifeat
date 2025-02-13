package database

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string
	Description string
	Image       string
	CategoryID  int32
	BrandID     int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SKU struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ProductID uuid.UUID
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Price struct {
	ID            int32 `gorm:"primaryKey;autoIncrement"`
	SkuID         uuid.UUID
	OriginalPrice int32
	EffectiveDate time.Time
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Inventory struct {
	ID           int32 `gorm:"primaryKey;autoIncrement"`
	SkuID        uuid.UUID
	Stock        int32
	Reservations json.RawMessage `gorm:"type:json"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SKUAttribute struct {
	ID          int32 `gorm:"primaryKey;autoIncrement"`
	SkuID       uuid.UUID
	AttributeID int32
	Value       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
