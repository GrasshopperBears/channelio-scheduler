package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func ConnectDatabase() {
	dsn := os.Getenv("DSN")
	openedDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	openedDb.AutoMigrate(&Schedule{})
	DB = openedDb

	log.Println("Database connected")
}

