package entity

import (
	"context"
	"errors"
	"time"

	"github.com/amirphl/matching-engine/market/domain/enum"
	coreEnum "github.com/amirphl/matching-engine/pkg/core/enum"
	"github.com/amirphl/matching-engine/pkg/db"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Market struct {
	db.UniversalCrossModel

	MarketPresenterId        string                       `gorm:"column:market_presenter_id;type:varchar(255);not null"`
	MarketName               string                       `gorm:"column:market_name;type:varchar(255);unique;not null"`
	MarketSymbol             string                       `gorm:"column:market_symbol;type:varchar(255);unique;not null"`
	MarketType               enum.StockMarketType         `gorm:"column:market_type;type:varchar(100);not null"`
	BaseAssetId              uuid.UUID                    `gorm:"column:base_asset_id;type:uuid;REFERENCES assets(reference_id);not null"`
	BaseAssetMinLotSize      decimal.Decimal              `gorm:"column:base_asset_min_lot_size;type:DECIMAL(25, 5);not null"`
	BaseAssetMaxLotSize      decimal.Decimal              `gorm:"column:base_asset_max_lot_size;type:DECIMAL(25, 5);not null"`
	QuoteAssetId             uuid.UUID                    `gorm:"column:quote_asset_id;type:uuid;;REFERENCES currencies(reference_id);not null"`
	QuoteAssetMinStepSize    decimal.Decimal              `gorm:"column:quote_asset_min_step_size;type:DECIMAL(25, 5);not null"`
	QuoteAssetMaxStepSize    decimal.Decimal              `gorm:"column:quote_asset_max_step_size;type:DECIMAL(25, 5);not null"`
	PriceStepSize            decimal.Decimal              `gorm:"column:price_step_size;type:DECIMAL(25, 5);not null"`
	QuantityStepSize         decimal.Decimal              `gorm:"column:quantity_step_size;type:DECIMAL(25, 5);not null"`
	PriceReference           decimal.Decimal              `gorm:"column:price_reference;type:DECIMAL(25, 5);not null"`
	PriceRangePercent        decimal.Decimal              `gorm:"column:price_range_percent;type:DECIMAL(25, 5);not null"`
	MakerCommissionPercent   decimal.Decimal              `gorm:"column:maker_commission_percent;type:DECIMAL(25, 5);not null"`
	MakerCommissionAssetType enum.CommissionAssetType     `gorm:"column:maker_commission_asset_type;type:varchar(100);not null"`
	MakerCommissionAssetId   uuid.UUID                    `gorm:"column:maker_commission_asset_id;not null"`
	TakerCommissionPercent   decimal.Decimal              `gorm:"column:taker_commission_percent;type:DECIMAL(25, 5);not null"`
	TakerCommissionAssetType enum.CommissionAssetType     `gorm:"column:taker_commission_asset_type;type:varchar(100);not null"`
	TakerCommissionAssetId   uuid.UUID                    `gorm:"column:taker_commission_asset_id;not null"`
	TaxPercent               decimal.Decimal              `gorm:"column:tax_percent;type:DECIMAL(25, 5);not null"`
	SelfTradePreventionType  enum.SelfTradePreventionType `gorm:"column:self_trade_prevention_type;type:varchar(100);not null" validate:"required"`
	OpenAt                   *int64                       `gorm:"column:open_at;type:timestamp"`
	CloseAt                  *int64                       `gorm:"column:close_at;type:timestamp"`
	StartAt                  *time.Time                   `gorm:"column:start_at;type:timestamp"`
	EndAt                    *time.Time                   `gorm:"column:end_at;type:timestamp"`
	StakeHoldersNotifyAt     *time.Time                   `gorm:"column:stake_holders_notify_at;type:timestamp"`
	MembersNotifyAt          *time.Time                   `gorm:"column:members_notify_at;type:timestamp"`
	Metadata                 *string                      `gorm:"column:metadata;type:jsonb"`
	Astat                    coreEnum.ActivityStatus      `gorm:"column:astat;type:varchar(100);not null"`

	// BaseAsset  assetEntity.Asset       `gorm:"constraint:OnDelete:CASCADE;foreignKey:base_asset_id;references:reference_id"`  // TODO: OnUpdate
	// QuoteAsset currencyEntity.Currency `gorm:"constraint:OnDelete:CASCADE;foreignKey:quote_asset_id;references:reference_id"` // TODO: OnUpdate
}

func (m Market) Validate(ctx context.Context) error {
	if m.ReferenceId == uuid.Nil {
		return errors.New("reference_id cannot be empty")
	}
	if m.MarketPresenterId == "" {
		return errors.New("market_presenter_id cannot be empty")
	}
	if m.MarketName == "" {
		return errors.New("market_name cannot be empty")
	}
	if m.MarketSymbol == "" {
		return errors.New("market_symbol cannot be empty")
	}
	if m.MarketType == "" {
		return errors.New("market_type cannot be empty")
	}
	if m.BaseAssetId == uuid.Nil {
		return errors.New("base_asset_id cannot be empty")
	}
	if m.QuoteAssetId == uuid.Nil {
		return errors.New("quote_asset_id cannot be empty")
	}
	if !m.PriceStepSize.IsPositive() {
		return errors.New("price_step_size must be greater than zero")
	}
	if !m.QuantityStepSize.IsPositive() {
		return errors.New("quantity_step_size must be greater than zero")
	}
	if !m.PriceReference.IsPositive() {
		return errors.New("price_reference must be greater than zero")
	}
	if m.PriceRangePercent.IsNegative() {
		return errors.New("price_range_percent cannot be negative")
	}
	if m.MakerCommissionPercent.IsNegative() {
		return errors.New("maker_commission_percent cannot be negative")
	}
	if m.TakerCommissionPercent.IsNegative() {
		return errors.New("taker_commission_percent cannot be negative")
	}
	if m.TaxPercent.IsNegative() {
		return errors.New("tax_percent cannot be negative")
	}
	if m.Astat == "" {
		return errors.New("astat cannot be empty")
	}
	if !m.BaseAssetMinLotSize.IsPositive() {
		return errors.New("base_asset_min_lot_size must be greater than 0")
	}
	if !m.BaseAssetMaxLotSize.IsPositive() {
		return errors.New("base_asset_max_lot_size must be greater than 0")
	}
	if !m.QuoteAssetMinStepSize.IsPositive() {
		return errors.New("quote_asset_min_step_size must be greater than 0")
	}
	if !m.QuoteAssetMaxStepSize.IsPositive() {
		return errors.New("quote_asset_max_step_size must be greater than 0")
	}

	return nil
}

type MarketFilter struct {
	Ids                    []int64
	ReferenceIds           []uuid.UUID
	MarketPresenterId      *string
	MarketName             *string
	MarketSymbol           *string
	MarketTypes            []enum.StockMarketType
	BaseAssetIds           []uuid.UUID
	QuoteAssetIds          []uuid.UUID
	BaseAssetMinLotSize    *decimal.Decimal
	BaseAssetMaxLotSize    *decimal.Decimal
	QuoteAssetMinStepSize  *decimal.Decimal
	QuoteAssetMaxStepSize  *decimal.Decimal
	PriceStepSize          *decimal.Decimal
	QuantityStepSize       *decimal.Decimal
	PriceReference         *decimal.Decimal
	PriceRangePercent      *decimal.Decimal
	MakerCommissionPercent *decimal.Decimal
	TakerCommissionPercent *decimal.Decimal
	TaxPercent             *decimal.Decimal
	Metadata               *string
	Astats                 []coreEnum.ActivityStatus

	db.Portion
}

func (m MarketFilter) Query(gd *gorm.DB) *gorm.DB {
	query := gd

	if len(m.Ids) > 0 {
		query = query.Where("id IN (?)", m.Ids)
	}
	if len(m.ReferenceIds) > 0 {
		query = query.Where("reference_id IN (?)", m.ReferenceIds)
	}
	if m.MarketPresenterId != nil {
		query = query.Where("market_presenter_id = ?", *m.MarketPresenterId)
	}
	if m.MarketName != nil {
		query = query.Where("market_name = ?", *m.MarketName)
	}
	if m.MarketSymbol != nil {
		query = query.Where("market_symbol = ?", *m.MarketSymbol)
	}
	if len(m.MarketTypes) > 0 {
		query = query.Where("market_type IN (?)", m.MarketTypes)
	}
	if len(m.BaseAssetIds) > 0 {
		query = query.Where("base_asset_id IN (?)", m.BaseAssetIds)
	}
	if len(m.QuoteAssetIds) > 0 {
		query = query.Where("quote_asset_id IN (?)", m.QuoteAssetIds)
	}
	if m.BaseAssetMinLotSize != nil {
		query = query.Where("base_asset_min_lot_size = ?", *m.BaseAssetMinLotSize)
	}
	if m.BaseAssetMaxLotSize != nil {
		query = query.Where("base_asset_max_lot_size = ?", *m.BaseAssetMaxLotSize)
	}
	if m.QuoteAssetMinStepSize != nil {
		query = query.Where("quote_asset_min_step_size = ?", *m.QuoteAssetMinStepSize)
	}
	if m.QuoteAssetMaxStepSize != nil {
		query = query.Where("quote_asset_max_step_size = ?", *m.QuoteAssetMaxStepSize)
	}
	if m.PriceStepSize != nil {
		query = query.Where("price_step_size = ?", *m.PriceStepSize)
	}
	if m.QuantityStepSize != nil {
		query = query.Where("quantity_step_size = ?", *m.QuantityStepSize)
	}
	if m.PriceReference != nil {
		query = query.Where("price_reference = ?", *m.PriceReference)
	}
	if m.PriceRangePercent != nil {
		query = query.Where("price_range_percent = ?", *m.PriceRangePercent)
	}
	if m.MakerCommissionPercent != nil {
		query = query.Where("maker_commission_percent = ?", *m.MakerCommissionPercent)
	}
	if m.TakerCommissionPercent != nil {
		query = query.Where("taker_commission_percent = ?", *m.TakerCommissionPercent)
	}
	if m.TaxPercent != nil {
		query = query.Where("tax_percent = ?", *m.TaxPercent)
	}
	if m.Metadata != nil {
		query = query.Where("metadata = ?", *m.Metadata)
	}
	if len(m.Astats) > 0 {
		query = query.Where("astat IN (?)", m.Astats)
	}

	return query
}
