package entity

import (
	"context"
	"errors"
	"time"

	"github.com/amirphl/matching-engine/market/domain/enum"
	"github.com/amirphl/matching-engine/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketLifecycle struct {
	db.UniversalModel

	MarketId         uuid.UUID                      `gorm:"column:market_id;type:uuid;REFERENCES markets(reference_id);not null"`
	Stage            enum.MarketLifecycleStage      `gorm:"column:stage;type:varchar(100);not null" validate:"required"`
	State            enum.MarketLifecycleState      `gorm:"column:state;type:varchar(100);not null" validate:"required"`
	Status           enum.MarketLifecycleStatus     `gorm:"column:status;type:varchar(100);not null" validate:"required"`
	TransitionAction enum.MarketLifecycleActionType `gorm:"column:transition_action;type:varchar(100);not null" validate:"required"`
	TransitionById   uuid.UUID                      `gorm:"column:transition_by_id;type:uuid;REFERENCES users(reference_id);not null"`
	Metadata         *string                        `gorm:"column:metadata;type:jsonb"`
	TimeoutAt        *time.Time                     `gorm:"column:timeout_at"`

	Market Market `gorm:"constraint:OnDelete:CASCADE;foreignKey:market_id;references:reference_id"` // TODO: OnUpdate
	// TransitionBy userEntity.User `gorm:"constraint:OnDelete:CASCADE;foreignKey:transition_by_id;references:reference_id"` // TODO: OnUpdate
}

func (m MarketLifecycle) Validate(ctx context.Context) error {
	if m.MarketId == uuid.Nil {
		return errors.New("market_id cannot be empty")
	}
	if m.Stage == "" {
		return errors.New("stage cannot be empty")
	}
	if m.State == "" {
		return errors.New("state cannot be empty")
	}
	if m.Status == "" {
		return errors.New("status cannot be empty")
	}
	if m.TransitionAction == "" {
		return errors.New("transition_action cannot be empty")
	}
	if m.TransitionById == uuid.Nil {
		return errors.New("transition_by_id cannot be empty")
	}

	return nil
}

type MarketLifecycleFilter struct {
	Ids               []uuid.UUID
	MarketIds         []uuid.UUID
	Stages            []enum.MarketLifecycleStage
	States            []enum.MarketLifecycleState
	Statuses          []enum.MarketLifecycleStatus
	TransitionActions []enum.MarketLifecycleActionType
	TransitionByIds   []uuid.UUID

	db.Portion
}

func (m MarketLifecycleFilter) Query(gd *gorm.DB) *gorm.DB {
	query := gd

	if len(m.Ids) > 0 {
		query = query.Where("id IN (?)", m.Ids)
	}
	if len(m.MarketIds) > 0 {
		query = query.Where("market_id IN (?)", m.MarketIds)
	}
	if len(m.Stages) > 0 {
		query = query.Where("stage IN (?)", m.Stages)
	}
	if len(m.States) > 0 {
		query = query.Where("state IN (?)", m.States)
	}
	if len(m.Statuses) > 0 {
		query = query.Where("status IN (?)", m.Statuses)
	}
	if len(m.TransitionActions) > 0 {
		query = query.Where("transition_action IN (?)", m.TransitionActions)
	}
	if len(m.TransitionByIds) > 0 {
		query = query.Where("transition_by_id IN (?)", m.TransitionByIds)
	}

	return query
}
