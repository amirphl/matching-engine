package db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrTransactionHasAlreadyStarted = errors.New("transaction has already been started")
)

type txKey struct{}

type UniversalModel struct {
	Id        uuid.UUID      `gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	Version   int64          `gorm:"column:version;not null;default:0"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;index"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type UniversalCrossModel struct {
	Id          int64           `gorm:"column:id;primary_key"`
	ReferenceId uuid.UUID       `gorm:"column:reference_id;type:uuid;unique"`
	Version     int64           `gorm:"column:version;not null;default:0"`
	CreatedAt   time.Time       `gorm:"column:created_at;not null;index"`
	UpdatedAt   *time.Time      `gorm:"column:updated_at;index"`
	DeletedAt   *gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type Portion struct {
	Offset int
	Limit  int
}

func GormConnection(ctx context.Context, gd *gorm.DB) *gorm.DB {
	tx, ok := gormTxFromContext(ctx)
	if ok {
		return tx
	}

	return gd
}

func gormTxFromContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	return tx, ok
}

func BeginTx(ctx context.Context, gd *gorm.DB) (*gorm.DB, context.Context, error) {
	tx, ok := gormTxFromContext(ctx)
	if ok {
		return tx, ctx, ErrTransactionHasAlreadyStarted
	}

	tx = gd.Begin()

	return tx, withTx(ctx, tx), nil
}

func withTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}
