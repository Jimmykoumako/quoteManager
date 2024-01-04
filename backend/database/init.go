package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "api/models"
)

var db *gorm.DB

// InitDB initializes the database connection
func InitDB() {
    connectionString := os.Getenv("DATABASE_URL")
    dbi, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to PostgreSQL")
    }
    db  = dbi

	fmt.Println("Connected to PostgreSQL...")

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