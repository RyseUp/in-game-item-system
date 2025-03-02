package main

import (
	"github.com/RyseUp/in-game-item-system/config"
	"github.com/RyseUp/in-game-item-system/internal/repositories"
	"github.com/RyseUp/in-game-item-system/internal/services"
	"log"
	"net/http"

	inventory "github.com/RyseUp/in-game-item-system/api/inventory/v1/v1connect"
	item "github.com/RyseUp/in-game-item-system/api/item/v1/v1connect"
	txn "github.com/RyseUp/in-game-item-system/api/transaction/v1/v1connect"
)

func main() {
	cfg := config.Load()
	config.ConnectDatabase(cfg.PostgresSQL)

	config.SeedDatabase()

	// item-service
	itemRepo := repositories.NewItemStore(config.DB)
	itemService := services.NewItemService(cfg, itemRepo)

	// inventory-service
	inventoryRepo := repositories.NewInventoryStore(config.DB)
	recordRepo := repositories.NewInventoryRecordStore(config.DB)
	inventoryService := services.NewInventoryAPI(cfg, inventoryRepo, recordRepo)

	// transaction-service
	transactionRepo := repositories.NewTransactionStore(config.DB)
	transactionService := services.NewTransactionAPI(cfg, transactionRepo)

	// transaction-consumer
	go services.ConsumeTransactionEvent(config.DB, transactionRepo)

	mux := http.NewServeMux()
	mux.Handle(item.NewItemServiceHandler(itemService))
	mux.Handle(inventory.NewInventoryAPIHandler(inventoryService))
	mux.Handle(txn.NewTransactionAPIHandler(transactionService))

	log.Println("Server running on :50051")
	http.ListenAndServe(":50051", mux)
}
