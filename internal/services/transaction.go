package services

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/RyseUp/in-game-item-system/api/transaction/v1"
	"github.com/RyseUp/in-game-item-system/api/transaction/v1/v1connect"
	"github.com/RyseUp/in-game-item-system/config"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"github.com/RyseUp/in-game-item-system/internal/repositories"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var (
	_ v1connect.TransactionAPIHandler = &TransactionAPI{}
)

type TransactionAPI struct {
	cfg             *config.Config
	transactionRepo repositories.Transaction
}

func NewTransactionAPI(
	cfg *config.Config,
	transactionRepo repositories.Transaction) *TransactionAPI {
	return &TransactionAPI{
		cfg:             cfg,
		transactionRepo: transactionRepo,
	}
}

func (s *TransactionAPI) CreateTransaction(ctx context.Context, c *connect.Request[v1.CreateTransactionRequest]) (*connect.Response[v1.CreateTransactionResponse], error) {
	var (
		req             = c.Msg
		userID          = req.GetUserId()
		itemID          = req.GetItemId()
		quantity        = req.GetQuantity()
		transactionType = req.GetTransactionType()
		preBalance      = req.GetPreBalance()
		postBalance     = req.GetPostBalance()
	)

	tx := s.transactionRepo.BeginTransaction()
	defer tx.Rollback()

	transaction := &models.Transaction{
		TransactionID:   uuid.New().String(),
		UserID:          userID,
		ItemID:          itemID,
		Quantity:        quantity,
		TransactionType: transactionType,
		PreBalance:      preBalance,
		PostBalance:     postBalance,
		CreatedAt:       time.Now(),
	}

	err := s.transactionRepo.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create transaction: %w", err))
	}

	tx.Commit()
	go PublishTransactionEvent(req.UserId, req.ItemId, transaction.TransactionID, req.Quantity, req.TransactionType, transaction.PreBalance, transaction.PostBalance)

	return connect.NewResponse(&v1.CreateTransactionResponse{Message: "Transaction recorded"}), nil
}

func (s *TransactionAPI) GetTransaction(ctx context.Context, c *connect.Request[v1.GetTransactionRequest]) (*connect.Response[v1.GetTransactionResponse], error) {
	var (
		req           = c.Msg
		transactionID = req.TransactionId
	)

	transaction, err := s.transactionRepo.GetTransactionByID(ctx, transactionID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, connect.NewError(connect.CodeNotFound, errors.New("transaction not found"))
	case err != nil:
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get transaction: %w", err))
	}

	return connect.NewResponse(&v1.GetTransactionResponse{
		TransactionId:   transaction.TransactionID,
		UserId:          transaction.UserID,
		ItemId:          transaction.ItemID,
		Quantity:        transaction.Quantity,
		TransactionType: transaction.TransactionType,
		PreBalance:      transaction.PreBalance,
		PostBalance:     transaction.PostBalance,
		CreatedAt:       transaction.CreatedAt.Format(time.RFC3339),
	}), nil
}

func (s *TransactionAPI) ListTransactions(ctx context.Context, c *connect.Request[v1.ListTransactionsRequest]) (*connect.Response[v1.ListTransactionsResponse], error) {
	var (
		req    = c.Msg
		userID = req.GetUserId()
		limit  = req.GetLimit()
		page   = req.GetPage()
	)

	transactions, err := s.transactionRepo.ListTransactions(ctx, userID, page, limit)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list transactions: %w", err))
	}

	var txnResponses []*v1.GetTransactionResponse
	for _, txn := range transactions {
		txnResponses = append(txnResponses, &v1.GetTransactionResponse{
			TransactionId:   txn.TransactionID,
			UserId:          txn.UserID,
			ItemId:          txn.ItemID,
			Quantity:        txn.Quantity,
			TransactionType: txn.TransactionType,
			PreBalance:      txn.PreBalance,
			PostBalance:     txn.PostBalance,
			CreatedAt:       txn.CreatedAt.Format(time.RFC3339),
		})
	}

	return connect.NewResponse(&v1.ListTransactionsResponse{Transactions: txnResponses}), nil
}
