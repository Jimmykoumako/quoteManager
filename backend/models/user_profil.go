package models

import (
	"time"
)

type UserProfile struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"userId" gorm:"unique;not null;index"` // Foreign key to User model
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Birthdate time.Time `json:"birthdate"`
    CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}