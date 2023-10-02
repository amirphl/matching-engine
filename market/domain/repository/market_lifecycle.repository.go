package repository

import (
	"context"

	"github.com/amirphl/matching-engine/market/domain/entity"
	"github.com/amirphl/matching-engine/pkg/db"
	"gorm.io/gorm"
)

type MarketLifecycleRepository interface {
	Create(ctx context.Context, in entity.MarketLifecycle) (res entity.MarketLifecycle, err error)
	Update(ctx context.Context, in entity.MarketLifecycle) (err error)
	Find(ctx context.Context, filter entity.MarketLifecycleFilter) (res []entity.MarketLifecycle, err error)
	FindById(ctx context.Context, id string) (res entity.MarketLifecycle, err error)
	FindByIds(ctx context.Context, ids []string) (res []entity.MarketLifecycle, err error)
	FindByMarketId(ctx context.Context, MarketId string) (res entity.MarketLifecycle, err error)
	Count(ctx context.Context, filter entity.MarketLifecycleFilter) (count int64, err error)
	Purge(ctx context.Context, id string) (err error)
	Delete(ctx context.Context, id string) (err error)
	FilterFind(ctx context.Context, query []any, order string, limit int, offset int) (res []entity.MarketLifecycle, err error)
	FilterCount(ctx context.Context, query []any) (count int64, err error)
	FindLast(ctx context.Context, filter entity.MarketLifecycleFilter) (res entity.MarketLifecycle, err error)
}

type marketLifecycleConfig struct {
	gd *gorm.DB
}

func NewMarketLifecycleRepository(gd *gorm.DB) MarketLifecycleRepository {
	return &marketLifecycleConfig{
		gd: gd,
	}
}

func (m marketLifecycleConfig) Create(ctx context.Context, in entity.MarketLifecycle) (res entity.MarketLifecycle, err error) {
	err = db.GormConnection(ctx, m.gd).Save(&in).Error
	if err != nil {
		return entity.MarketLifecycle{}, err
	}

	return in, nil
}

func (m marketLifecycleConfig) Update(ctx context.Context, in entity.MarketLifecycle) (err error) {
	err = db.GormConnection(ctx, m.gd).Save(&in).Error
	if err != nil {
		return err
	}

	return nil
}

func (m marketLifecycleConfig) Find(ctx context.Context, filter entity.MarketLifecycleFilter) (res []entity.MarketLifecycle, err error) {
	query := filter.Query(m.gd)
	err = db.GormConnection(ctx, m.gd).Model(&res).Limit(filter.Limit).Offset(filter.Offset).Find(&res, query).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m marketLifecycleConfig) FindById(ctx context.Context, id string) (res entity.MarketLifecycle, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).Last(&res, "id = ?", id).Error
	if err != nil {
		return entity.MarketLifecycle{}, err
	}

	return res, nil
}

func (m marketLifecycleConfig) FindByIds(ctx context.Context, ids []string) (res []entity.MarketLifecycle, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).Find(&res, "id IN (?)", ids).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m marketLifecycleConfig) FindByMarketId(ctx context.Context, MarketId string) (res entity.MarketLifecycle, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).Last(&res, "market_id = ?", MarketId).Error
	if err != nil {
		return entity.MarketLifecycle{}, err
	}

	return res, nil
}

func (m marketLifecycleConfig) Count(ctx context.Context, filter entity.MarketLifecycleFilter) (count int64, err error) {
	query := filter.Query(m.gd)
	err = db.GormConnection(ctx, m.gd).Model(&entity.MarketLifecycle{}).Where(query).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m marketLifecycleConfig) Purge(ctx context.Context, id string) (err error) {
	err = db.GormConnection(ctx, m.gd).Exec("DELETE FROM market_lifecycles WHERE id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (m marketLifecycleConfig) Delete(ctx context.Context, id string) (err error) {
	err = db.GormConnection(ctx, m.gd).Delete(&entity.MarketLifecycle{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (m marketLifecycleConfig) FilterFind(ctx context.Context, query []any, order string, limit int, offset int) (res []entity.MarketLifecycle, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&res, query...).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m marketLifecycleConfig) FilterCount(ctx context.Context, query []any) (count int64, err error) {
	countQuery := db.GormConnection(ctx, m.gd).Model(&entity.MarketLifecycle{})
	if len(query) > 1 {
		countQuery = countQuery.Where(query[0], query[1:]...)
	}

	err = countQuery.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m marketLifecycleConfig) FindLast(ctx context.Context, filter entity.MarketLifecycleFilter) (res entity.MarketLifecycle, err error) {
	query := filter.Query(m.gd)
	err = db.GormConnection(ctx, m.gd).Model(&res).Order("created_at desc").Last(&res, query).Error
	if err != nil {
		return entity.MarketLifecycle{}, err
	}

	return res, nil
}
