package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"api/models"
)

var db *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	database, err := gorm.Open(postgres.Open("user=username dbname=mydatabase sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db = database

	// Auto-migrate the schema
	db.AutoMigrate(&models.Quote{}, &models.Feedback{}, &models.User{}, &models.Category{}, &models.Tag{}, &models.Like{}, &models.LiteraryWork{}, &models.Folder{})
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database connection")
	}

	sqlDB.Close()
}
