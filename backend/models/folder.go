package models

import "time"

// Folder Entity
type Folder struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	UserID    uint      `json:"userId" gorm:"not null"`
	Quotes    []Quote   `json:"quotes" gorm:"many2many:quote_folders;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
