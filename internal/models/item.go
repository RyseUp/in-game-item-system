package models

import (
	"database/sql"
	"time"
)

type Item struct {
	ID          int64        `gorm:"type:bigint;autoIncrement"`
	ItemID      string       `gorm:"type:text;primaryKey" json:"item_id"`
	Name        string       `gorm:"type:text;not null;unique;index"`
	Description string       `gorm:"type:text" json:"description"`
	CreatedAt   time.Time    `gorm:"column:created_at;type:timestamptz;default:now()" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;type:timestamptz;default:now()" json:"updated_at"`
	DeletedAt   sql.NullTime `gorm:"column:deleted_at;type:timestamptz" json:"deleted_at"`
}

func (Item) TableName() string {
	return "game_items"
}
