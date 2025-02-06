package repo

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type SQLStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) Querier {
	return &SQLStore{
		db: db,
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(Querier) error) error {
	// Bắt đầu transaction với GORM
	tx := store.db.WithContext(ctx).Begin()

	fmt.Println("Start create user")

	if tx.Error != nil {
		return tx.Error
	}

	q := New(tx)
	// Thực thi function với transaction
	err := fn(q)
	if err != nil {
		// Nếu có lỗi, rollback transaction
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	fmt.Println("Done create user")

	// Commit transaction
	return tx.Commit().Error
}
