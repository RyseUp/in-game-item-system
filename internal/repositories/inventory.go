package repositories

import (
	"context"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Inventory interface {
	GetInventoryByUserIDAndItemID(ctx context.Context, userID, itemID string) (*models.Inventory, error)
	UpdateInventoryQuantity(ctx context.Context, tx *gorm.DB, inventory *models.Inventory) error
	CreateInventory(ctx context.Context, tx *gorm.DB, inventory *models.Inventory) error
	BeginTransaction() *gorm.DB
}

var _ Inventory = &InventoryStore{}

type InventoryStore struct {
	db *gorm.DB
}

func NewInventoryStore(db *gorm.DB) *InventoryStore {
	return &InventoryStore{db: db}
}

func (s *InventoryStore) GetInventoryByUserIDAndItemID(ctx context.Context, userID, itemID string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := s.db.WithContext(ctx).
		Model(&models.Inventory{}).
		Where("user_id = ? AND item_id = ?", userID, itemID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&inventory).
		Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (s *InventoryStore) UpdateInventoryQuantity(ctx context.Context, tx *gorm.DB, inventory *models.Inventory) error {
	return tx.WithContext(ctx).
		Model(&models.Inventory{}).
		Where("user_id = ? AND item_id = ?", inventory.UserID, inventory.ItemID).
		Update("quantity", inventory.Quantity).
		Error
}

func (s *InventoryStore) CreateInventory(ctx context.Context, tx *gorm.DB, inventory *models.Inventory) error {
	return tx.WithContext(ctx).Create(inventory).Error
}

func (s *InventoryStore) BeginTransaction() *gorm.DB {
	return s.db.Begin()
}
