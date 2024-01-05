// migrations/001_create_quotes_table.go
package migrations

import (
	"api/models" // Import your models package
	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Quote{})
}

func Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Quote{})
}
