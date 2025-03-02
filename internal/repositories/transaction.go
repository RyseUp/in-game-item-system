package repositories

import (
	"context"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"gorm.io/gorm"
)

type Transaction interface {
	CreateTransaction(ctx context.Context, tx *gorm.DB, transaction *models.Transaction) error
	GetTransactionByID(ctx context.Context, transactionID string) (*models.Transaction, error)
	ListTransactions(ctx context.Context, userID string, page, limit int32) ([]*models.Transaction, error)
	BeginTransaction() *gorm.DB
}

var _ Transaction = &TransactionStore{}

type TransactionStore struct {
	db *gorm.DB
}

func NewTransactionStore(db *gorm.DB) *TransactionStore {
	return &TransactionStore{db: db}
}

func (s *TransactionStore) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction *models.Transaction) error {
	return tx.WithContext(ctx).Model(&models.Transaction{}).Create(transaction).Error
}

func (s *TransactionStore) GetTransactionByID(ctx context.Context, transactionID string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := s.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("transaction_id = ?", transactionID).
		First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionStore) ListTransactions(ctx context.Context, userID string, page, limit int32) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := s.db.WithContext(ctx).
		Model(&models.Transaction{}).
		Where("user_id = ?", userID).
		Limit(int(limit)).
		Offset(int((page - 1) * limit)).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *TransactionStore) BeginTransaction() *gorm.DB {
	return s.db.Begin()
}
