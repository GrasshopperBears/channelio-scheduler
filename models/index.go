package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDatabase() {
	dsn := os.Getenv("DSN")

	if _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected")
}

