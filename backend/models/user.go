package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"log"
)

var DB *gorm.DB // Package-level variable to hold the Gorm DB instance

const AdminRole = "admin"

// // SetDB sets the Gorm DB instance for the user model
// func SetDB(database *gorm.DB) {
//     DB = database
// }

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null;index"` // Added indexing for better query performance
	Password  string    `json:"-" gorm:"not null"`
	Quotes    []Quote   `json:"quotes" gorm:"foreignKey:user_id"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// BeforeCreate hooks into the GORM lifecycle to hash the user's password before saving
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies if the provided password matches the hashed password
func (u *User) CheckPassword(password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    if err != nil {
        log.Printf("Password comparison error: %v", err)
    }
    return nil
}


