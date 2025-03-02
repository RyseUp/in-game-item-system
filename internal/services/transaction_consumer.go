package services

import (
	"context"
	"encoding/json"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"github.com/RyseUp/in-game-item-system/internal/repositories"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"log"
)

type TransactionEvent struct {
	TransactionID   string `json:"transaction_id"`
	UserID          string `json:"user_id"`
	ItemID          string `json:"item_id"`
	Quantity        int32  `json:"quantity"`
	TransactionType string `json:"transaction_type"`
}

func ConsumeTransactionEvent(db *gorm.DB, transactionRepo repositories.Transaction) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"transaction_events",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	for msg := range msgs {
		var event TransactionEvent
		err := json.Unmarshal(msg.Body, &event)
		if err != nil {
			log.Printf("Failed to decode event: %v", err)
			continue
		}

		log.Printf("Processing transaction: %+v", event)

		tx := db.Begin()
		defer tx.Rollback()

		transaction := &models.Transaction{
			TransactionID:   event.TransactionID,
			UserID:          event.UserID,
			ItemID:          event.ItemID,
			Quantity:        event.Quantity,
			TransactionType: event.TransactionType,
		}

		err = transactionRepo.CreateTransaction(context.Background(), tx, transaction)
		if err != nil {
			log.Printf("Failed to store transaction: %v", err)
			continue
		}

		tx.Commit()
		log.Printf("Stored transaction successfully: %s", event.TransactionID)
	}
}
