package database

import (
	"api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"api/logger"
)

var db *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	connectionString := os.Getenv("DATABASE_URL")
	dbi, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL")
	}
	db = dbi

	logger.Log.Info("Connected to PostgreSQL...")

	// Auto-migrate the schema
	db.AutoMigrate(&models.Like{}, &models.Folder{}, &models.LiteraryWork{}, &models.Feedback{}, &models.Tag{}, &models.Category{}, &models.User{}, &models.Quote{})
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database connection")
	}

	sqlDB.Close()
}

// GetDB returns the reference to the database instance
func GetDB() *gorm.DB {
	return db
}
