package models

import (
	"time"
)

type Transaction struct {
	ID              int64     `gorm:"type:bigint;autoIncrement"`
	TransactionID   string    `gorm:"type:text;primaryKey" json:"transaction_id"`
	UserID          string    `gorm:"type:text;not null;index" json:"user_id"`
	ItemID          string    `gorm:"type:text;not null;index" json:"item_id"`
	Quantity        int32     `gorm:"type:int;not null" json:"quantity"`
	TransactionType string    `gorm:"type:text;not null" json:"transaction_type"` // e.g., "PURCHASE", "USE"
	PreBalance      int32     `gorm:"type:int;not null" json:"pre_balance"`
	PostBalance     int32     `gorm:"type:int;not null" json:"post_balance"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamptz;default:now()" json:"created_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
