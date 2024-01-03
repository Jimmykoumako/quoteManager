package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetDB sets the Gorm DB instance for the models package
func SetDB(database *gorm.DB) {
	DB = database
}
