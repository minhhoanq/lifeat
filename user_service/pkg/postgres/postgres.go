package postgres

import (
	"context"
	"fmt"

	"github.com/minhhoanq/lifeat/common/logger"
	"github.com/minhhoanq/lifeat/user_service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	*gorm.DB
}

func New(cfg config.Config, l logger.Interface) (Postgres, error) {
	dbMigration, err := NewMigrator(cfg)
	if err != nil {
		return Postgres{}, err
	}

	err = dbMigration.Up(context.Background())
	if err != nil {
		return Postgres{}, err
	}

	db, err := NewDatabase(cfg, l)
	if err != nil {
		return Postgres{}, nil
	}

	return Postgres{DB: db}, nil
}

func NewDatabase(cfg config.Config, l logger.Interface) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	l.Info("connection to datatbase successfully")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *Postgres) Close(l logger.Interface) {
	db, _ := p.DB.DB()
	db.Close()
	l.Info("disconnect to datatbase")
}
