package repositories

import (
	"context"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"gorm.io/gorm"
)

type InventoryRecord interface {
	CreateRecord(ctx context.Context, tx *gorm.DB, record *models.InventoryRecord) error
}

var _ InventoryRecord = &InventoryRecordStore{}

type InventoryRecordStore struct {
	db *gorm.DB
}

func NewInventoryRecordStore(db *gorm.DB) *InventoryRecordStore {
	return &InventoryRecordStore{db: db}
}

func (s *InventoryRecordStore) CreateRecord(ctx context.Context, tx *gorm.DB, record *models.InventoryRecord) error {
	return tx.WithContext(ctx).Create(record).Error
}
