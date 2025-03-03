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

func (s *InventoryAPI) UserGetInventory(ctx context.Context, c *connect.Request[v1.UserGetInventoryRequest]) (*connect.Response[v1.UserGetInventoryResponse], error) {
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

	return connect.NewResponse(&v1.UserGetInventoryResponse{
		UserId:   inventory.UserID,
		ItemId:   inventory.ItemID,
		Quantity: inventory.Quantity,
	}), nil
}

func (s *InventoryAPI) UserAddItemInInventory(ctx context.Context, c *connect.Request[v1.UserAddItemInInventoryRequest]) (*connect.Response[v1.UserAddItemInInventoryResponse], error) {
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

	go PublishTransactionEvent(userID, itemID, record.RecordID, quantity, "INVENTORY_UPDATE", record.PreBalance, record.PostBalance)
	return connect.NewResponse(&v1.UserAddItemInInventoryResponse{
		Message: "success",
	}), nil
}

func (s *InventoryAPI) UserUseItemInInventory(ctx context.Context, c *connect.Request[v1.UserUseItemInInventoryRequest]) (*connect.Response[v1.UserUseItemInInventoryResponse], error) {
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

	if inventory.Quantity < quantity {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("insufficient item quantity"))
	}

	preBalance := inventory.Quantity
	inventory.Quantity -= quantity

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

	go PublishTransactionEvent(userID, itemID, record.RecordID, quantity, "ITEM_USED", record.PreBalance, record.PostBalance)

	return connect.NewResponse(&v1.UserUseItemInInventoryResponse{
		Message: "Item used successfully!",
	}), nil
}
