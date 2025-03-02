package models

import (
	"time"
)

type InventoryRecord struct {
	ID          int64     `gorm:"type:bigint;autoIncrement"`
	RecordID    string    `gorm:"type:text;primaryKey" json:"history_id"`
	UserID      string    `gorm:"type:text;not null;index" json:"user_id"`
	ItemID      string    `gorm:"type:text;not null;index" json:"item_id"`
	PreBalance  int32     `gorm:"type:int;not null" json:"pre_balance"`
	PostBalance int32     `gorm:"type:int;not null" json:"post_balance"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamptz;default:now()" json:"updated_at"`
}

func (InventoryRecord) TableName() string {
	return "inventory_records"
}
