package models

import "time"

// LiteraryWork Entity
type LiteraryWork struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Author    string    `json:"author"`
	Quotes    []Quote   `json:"quotes"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}