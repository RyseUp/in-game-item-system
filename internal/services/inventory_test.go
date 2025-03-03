package services

import (
	"context"
	"testing"

	"github.com/RyseUp/in-game-item-system/api/inventory/v1"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"github.com/RyseUp/in-game-item-system/internal/repositories/mocks"
	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupMockInventoryService() (*InventoryAPI, *mocks.Inventory) {
	mockInventoryRepo := new(mocks.Inventory)
	mockInventoryRecordRepo := new(mocks.InventoryRecord)
	service := NewInventoryAPI(nil, mockInventoryRepo, mockInventoryRecordRepo)
	return service, mockInventoryRepo
}

func TestUserGetInventory(t *testing.T) {
	service, mockRepo := setupMockInventoryService()

	mockInventory := &models.Inventory{
		UserID:   "user_001",
		ItemID:   "item_001",
		Quantity: 5,
	}

	mockRepo.On("GetInventoryByUserIDAndItemID", mock.Anything, "user_001", "item_001").
		Return(mockInventory, nil)

	req := &v1.UserGetInventoryRequest{
		UserId: "user_001",
		ItemId: "item_001",
	}

	resp, err := service.UserGetInventory(context.Background(), connect.NewRequest(req))
	assert.NoError(t, err)
	assert.Equal(t, int32(5), resp.Msg.Quantity)

	mockRepo.AssertExpectations(t)
}
