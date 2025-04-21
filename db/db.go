package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"krishimitra-api/models"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	migErr := database.AutoMigrate(&models.User{}, &models.Crop{}, &models.Product{}, &models.NewsArticle{}, &models.HistoryItem{})
	if migErr != nil {
		log.Fatal("Failed to migrate the database:", migErr)
	}
	DB = database
}
