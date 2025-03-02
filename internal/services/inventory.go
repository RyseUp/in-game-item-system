package services

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/RyseUp/in-game-item-system/api/inventory/v1"
	"github.com/RyseUp/in-game-item-system/api/inventory/v1/v1connect"
	"github.com/RyseUp/in-game-item-system/config"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"github.com/RyseUp/in-game-item-system/internal/repositories"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	_ v1connect.InventoryAPIHandler = &InventoryAPI{}
)

type InventoryAPI struct {
	cfg                 *config.Config
	inventoryRepo       repositories.Inventory
	inventoryRecordRepo repositories.InventoryRecord
}

func NewInventoryAPI(
	cfg *config.Config,
	inventoryRepo repositories.Inventory,
	inventoryRecordRepo repositories.InventoryRecord) *InventoryAPI {
	return &InventoryAPI{
		cfg:                 cfg,
		inventoryRepo:       inventoryRepo,
		inventoryRecordRepo: inventoryRecordRepo,
	}
}

func (s *InventoryAPI) GetInventory(ctx context.Context, c *connect.Request[v1.GetInventoryRequest]) (*connect.Response[v1.GetInventoryResponse], error) {
	var (
		req    = c.Msg
		userID = req.GetUserId()
		itemID = req.GetItemId()
	)

	inventory, err := s.inventoryRepo.GetInventoryByUserIDAndItemID(ctx, userID, itemID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("failed to get item: %w", err))
	case err != nil:
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get item: %w", err))
	}

	return connect.NewResponse(&v1.GetInventoryResponse{
		UserId:   inventory.UserID,
		ItemId:   inventory.ItemID,
		Quantity: inventory.Quantity,
	}), nil
}

func (s *InventoryAPI) UpdateInventory(ctx context.Context, c *connect.Request[v1.UpdateInventoryRequest]) (*connect.Response[v1.UpdateInventoryResponse], error) {
	var (
		req      = c.Msg
		userID   = req.GetUserId()
		itemID   = req.GetItemId()
		quantity = req.GetQuantity()
	)

	tx := s.inventoryRepo.BeginTransaction()
	defer tx.Rollback()

	inventory, err := s.inventoryRepo.GetInventoryByUserIDAndItemID(ctx, userID, itemID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("inventory record not found"))
	} else if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to retrieve inventory: %w", err))
	}

	preBalance := inventory.Quantity
	inventory.Quantity += quantity

	err = s.inventoryRepo.UpdateInventoryQuantity(ctx, tx, inventory)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update inventory: %w", err))
	}

	record := &models.InventoryRecord{
		RecordID:    uuid.New().String(),
		UserID:      userID,
		ItemID:      itemID,
		PreBalance:  preBalance,
		PostBalance: inventory.Quantity,
	}
	err = s.inventoryRecordRepo.CreateRecord(ctx, tx, record)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to record inventory update: %w", err))
	}

	tx.Commit()

	go PublishTransactionEvent(userID, itemID, record.RecordID, quantity, "INVENTORY_UPDATE")
	return connect.NewResponse(&v1.UpdateInventoryResponse{
		Message: "success",
	}), nil
}
