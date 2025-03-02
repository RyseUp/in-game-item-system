package models

import (
	"time"
)

type Inventory struct {
	ID          int64     `gorm:"type:bigint;autoIncrement"`
	InventoryID string    `gorm:"type:text;primaryKey" json:"inventory_id"`
	UserID      string    `gorm:"type:text;not null;index" json:"user_id"`
	ItemID      string    `gorm:"type:text;not null;index" json:"item_id"`
	Quantity    int32     `gorm:"type:int;not null;default:0" json:"quantity"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamptz;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamptz;default:now()" json:"updated_at"`
}

func (Inventory) TableName() string {
	return "inventory"
}
