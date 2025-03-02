package config

import (
	"fmt"
	"github.com/RyseUp/in-game-item-system/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase(dsn string) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbName, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.Item{},
		&models.Inventory{},
		&models.InventoryRecord{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}

	log.Println("Successfully connected to database")
	DB = db
}
