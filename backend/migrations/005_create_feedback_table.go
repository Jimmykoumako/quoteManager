// migrations/001_create_quotes_table.go
package migrations

import (
	"gorm.io/gorm"
	"api/models" // Import your models package
)

func Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Feedback{})
}

func Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Feedback{})
}
