package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDatabase() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	db.AutoMigrate(&Schedule{})

	log.Println("Database connected")
}

