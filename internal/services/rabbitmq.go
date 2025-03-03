package services

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func getRabbitMQURL() string {
	if os.Getenv("DOCKER_ENV") == "true" {
		return "amqp://guest:guest@rabbitmq:5672/"
	}
	return "amqp://guest:guest@rabbitmq:5672/"
}

func PublishTransactionEvent(userID, itemID, transactionID string, quantity int32, txnType string, preBalance, postBalance int32) {
	rabbitMQURL := getRabbitMQURL()
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return
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
		log.Printf("Failed to declare a queue: %v", err)
	}

	body := fmt.Sprintf(`{"transaction_id": "%s", "user_id": "%s", "item_id": "%s", "quantity": %d, "transaction_type": "%s", "pre_balance": %d, "post_balance": %d}`,
		transactionID, userID, itemID, quantity, txnType, preBalance, postBalance)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return
	}
	log.Printf("Published transaction event: %s", body)
}
