package repositories

import (
	"context"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"gorm.io/gorm"
	"time"
)

type Item interface {
	GetItemByID(ctx context.Context, itemID string) (*models.Item, error)
	CreateItem(ctx context.Context, item *models.Item) error
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, itemID string) error
	ListItem(ctx context.Context, limit, offset int32) ([]*models.Item, error)
}

var _ Item = &ItemStore{}

type ItemStore struct {
	db *gorm.DB
}

func NewItemStore(db *gorm.DB) *ItemStore {
	return &ItemStore{db: db}
}

func (s *ItemStore) GetItemByID(ctx context.Context, itemID string) (*models.Item, error) {
	var item models.Item
	err := s.db.WithContext(ctx).
		Model(&models.Item{}).
		Where("item_id = ?", itemID).
		First(&item).
		Error

	if err != nil {
		return nil, err
	}

	if !item.DeletedAt.Time.IsZero() {
		return nil, gorm.ErrRecordNotFound
	}

	return &item, err
}

func (s *ItemStore) CreateItem(ctx context.Context, item *models.Item) error {
	return s.db.WithContext(ctx).Create(item).Error
}

func (s *ItemStore) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.db.WithContext(ctx).
		Model(&models.Item{}).
		Where("item_id = ?", item.ItemID).
		Updates(item).
		Error
}

func (s *ItemStore) DeleteItem(ctx context.Context, itemID string) error {
	return s.db.WithContext(ctx).
		Model(&models.Item{}).
		Where("item_id = ?", itemID).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
		}).
		Error
}

func (s *ItemStore) ListItem(ctx context.Context, page, limit int32) ([]*models.Item, error) {
	var items []*models.Item
	err := s.db.WithContext(ctx).
		Limit(int(limit)).
		Offset(int((page - 1) * limit)).
		Where("deleted_at is null").
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}
	return items, nil
}
