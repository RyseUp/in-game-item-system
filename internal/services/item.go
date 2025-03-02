package services

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/RyseUp/in-game-item-system/api/item/v1"
	"github.com/RyseUp/in-game-item-system/api/item/v1/v1connect"
	"github.com/RyseUp/in-game-item-system/config"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"github.com/RyseUp/in-game-item-system/internal/repositories"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	_ v1connect.ItemServiceHandler = &ItemAPI{}
)

type ItemAPI struct {
	cfg      *config.Config
	itemRepo repositories.Item
}

func NewItemService(
	cfg *config.Config,
	itemRepo repositories.Item) *ItemAPI {
	return &ItemAPI{
		cfg:      cfg,
		itemRepo: itemRepo,
	}
}

func (s *ItemAPI) GetItem(ctx context.Context, c *connect.Request[v1.GetItemRequest]) (*connect.Response[v1.GetItemResponse], error) {
	var (
		req    = c.Msg
		itemID = req.GetItemId()
	)

	item, err := s.itemRepo.GetItemByID(ctx, itemID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, connect.NewError(connect.CodeNotFound, errors.New("item not found"))
	case err != nil:
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get item: %w", err))
	}

	return connect.NewResponse(&v1.GetItemResponse{
		ItemId:      item.ItemID,
		Name:        item.Name,
		Description: item.Description,
	}), nil
}

func (s *ItemAPI) CreateItem(ctx context.Context, c *connect.Request[v1.CreateItemRequest]) (*connect.Response[v1.CreateItemResponse], error) {
	var (
		req             = c.Msg
		itemName        = req.GetName()
		itemDescription = req.GetDescription()
	)

	itemID, err := uuid.NewUUID()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create item: %w", err))
	}

	newItem := &models.Item{
		ItemID:      itemID.String(),
		Name:        itemName,
		Description: itemDescription,
	}

	err = s.itemRepo.CreateItem(ctx, newItem)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create item: %w", err))
	}

	return connect.NewResponse(&v1.CreateItemResponse{
		ItemId: newItem.ItemID,
	}), nil
}

func (s *ItemAPI) UpdateItem(ctx context.Context, c *connect.Request[v1.UpdateItemRequest]) (*connect.Response[v1.UpdateItemResponse], error) {
	var (
		req    = c.Msg
		itemID = req.GetItemId()
	)

	existingItem, err := s.itemRepo.GetItemByID(ctx, itemID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, connect.NewError(connect.CodeNotFound, errors.New("item not found"))
	case err != nil:
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get item: %w", err))
	}

	existingItem.Name = req.GetName()
	existingItem.Description = req.GetDescription()

	err = s.itemRepo.UpdateItem(ctx, existingItem)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update item: %w", err))
	}

	return connect.NewResponse(&v1.UpdateItemResponse{
		Message: "success",
	}), nil
}

func (s *ItemAPI) DeleteItem(ctx context.Context, c *connect.Request[v1.DeleteItemRequest]) (*connect.Response[v1.DeleteItemResponse], error) {
	var (
		req = c.Msg
	)

	_, err := s.itemRepo.GetItemByID(ctx, req.GetItemId())
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, connect.NewError(connect.CodeNotFound, errors.New("item not found"))
	case err != nil:
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get item: %w", err))
	}

	err = s.itemRepo.DeleteItem(ctx, req.GetItemId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to delete item: %w", err))
	}

	return connect.NewResponse(&v1.DeleteItemResponse{
		Message: "deleted",
	}), nil
}

func (s *ItemAPI) ListItems(ctx context.Context, c *connect.Request[v1.ListItemsRequest]) (*connect.Response[v1.ListItemsResponse], error) {
	var (
		req   = c.Msg
		page  = req.GetPage()
		limit = req.GetLimit()
	)

	items, err := s.itemRepo.ListItem(ctx, page, limit)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list items: %w", err))
	}

	listItems := make([]*v1.Item, len(items))
	for i, item := range items {
		listItems[i] = &v1.Item{
			ItemId:      item.ItemID,
			Name:        item.Name,
			Description: item.Description,
		}
	}

	return connect.NewResponse(&v1.ListItemsResponse{
		Items: listItems,
	}), nil
}
