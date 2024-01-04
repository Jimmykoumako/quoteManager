package models

import (
	"gorm.io/gorm"
)


// SetDB sets the Gorm DB instance for the models package
func SetDB(database *gorm.DB) {
	DB = database
}
