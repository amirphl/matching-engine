package repository

import (
	"context"

	"github.com/amirphl/matching-engine/market/domain/entity"
	"github.com/amirphl/matching-engine/pkg/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MarketRepository interface {
	Create(ctx context.Context, in entity.Market) (res entity.Market, err error)
	Update(ctx context.Context, in entity.Market) (err error)
	Find(ctx context.Context, filter entity.MarketFilter) (res []entity.Market, err error)
	FindById(ctx context.Context, id int64) (res entity.Market, err error)
	FindByIds(ctx context.Context, ids []int64) (res []entity.Market, err error)
	FindByReferenceId(ctx context.Context, referenceId string) (res entity.Market, err error)
	Count(ctx context.Context, filter entity.MarketFilter) (count int64, err error)
	Purge(ctx context.Context, id int64) (err error)
	Delete(ctx context.Context, id int64) (err error)
	FilterFind(ctx context.Context, query []any, order string, limit int, offset int) (res []entity.Market, err error)
	FilterCount(ctx context.Context, query []any) (count int64, err error)
}

type marketConfig struct {
	gd *gorm.DB
}

func NewMarketRepository(gd *gorm.DB) MarketRepository {
	return &marketConfig{
		gd: gd,
	}
}

func (m marketConfig) Create(ctx context.Context, in entity.Market) (res entity.Market, err error) {
	err = db.GormConnection(ctx, m.gd).Save(&in).Error
	if err != nil {
		return entity.Market{}, err
	}

	return in, nil
}

func (m marketConfig) Update(ctx context.Context, in entity.Market) (err error) {
	err = db.GormConnection(ctx, m.gd).Save(&in).Error
	if err != nil {
		return err
	}

	return nil
}

func (m marketConfig) Find(ctx context.Context, filter entity.MarketFilter) (res []entity.Market, err error) {
	query := filter.Query(m.gd)
	err = db.GormConnection(ctx, m.gd).Model(&res).Limit(filter.Limit).Offset(filter.Offset).Find(&res, query).Error
	if err != nil {
		return nil, err
	}

	res = fixPercents(res)

	return res, nil
}

func (m marketConfig) FindById(ctx context.Context, id int64) (res entity.Market, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).Last(&res, "id = ?", id).Error
	if err != nil {
		return entity.Market{}, err
	}

	res = fixPercent(res)

	return res, nil
}

func (m marketConfig) FindByIds(ctx context.Context, ids []int64) (res []entity.Market, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).Find(&res, "id IN (?)", ids).Error
	if err != nil {
		return nil, err
	}

	res = fixPercents(res)

	return res, nil
}

func (m marketConfig) FindByReferenceId(ctx context.Context, referenceId string) (res entity.Market, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).Last(&res, "reference_id = ?", referenceId).Error
	if err != nil {
		return entity.Market{}, err
	}

	res = fixPercent(res)

	return res, nil
}

func (m marketConfig) Count(ctx context.Context, filter entity.MarketFilter) (count int64, err error) {
	query := filter.Query(m.gd)
	err = db.GormConnection(ctx, m.gd).Model(&entity.Market{}).Where(query).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m marketConfig) Purge(ctx context.Context, id int64) (err error) {
	err = db.GormConnection(ctx, m.gd).Exec("DELETE FROM markets WHERE id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (m marketConfig) Delete(ctx context.Context, id int64) (err error) {
	err = db.GormConnection(ctx, m.gd).Delete(&entity.Market{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (m marketConfig) FilterFind(ctx context.Context, query []any, order string, limit int, offset int) (res []entity.Market, err error) {
	err = db.GormConnection(ctx, m.gd).Model(&res).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&res, query...).Error
	if err != nil {
		return nil, err
	}

	res = fixPercents(res)

	return res, nil
}

func (m marketConfig) FilterCount(ctx context.Context, query []any) (count int64, err error) {
	countQuery := db.GormConnection(ctx, m.gd).Model(&entity.Market{})
	if len(query) > 1 {
		countQuery = countQuery.Where(query[0], query[1:]...)
	}

	err = countQuery.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FIXME:
func fixPercents(res []entity.Market) []entity.Market {
	for idx := range res {
		res[idx].MakerCommissionPercent = res[idx].MakerCommissionPercent.Div(decimal.NewFromFloat(100))
		res[idx].TakerCommissionPercent = res[idx].TakerCommissionPercent.Div(decimal.NewFromFloat(100))
		res[idx].PriceRangePercent = res[idx].PriceRangePercent.Div(decimal.NewFromFloat(100))
		res[idx].TaxPercent = res[idx].TaxPercent.Div(decimal.NewFromFloat(100))
	}

	return res
}

// FIXME:
func fixPercent(res entity.Market) entity.Market {
	res.MakerCommissionPercent = res.MakerCommissionPercent.Div(decimal.NewFromFloat(100))
	res.TakerCommissionPercent = res.TakerCommissionPercent.Div(decimal.NewFromFloat(100))
	res.PriceRangePercent = res.PriceRangePercent.Div(decimal.NewFromFloat(100))
	res.TaxPercent = res.TaxPercent.Div(decimal.NewFromFloat(100))

	return res
}
