package models

import "time"

// Feedback Entity
type Feedback struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	Rating    int       `json:"rating" gorm:"not null;check:(rating >= 1) AND (rating <= 5)"`
	QuoteID   uint      `json:"quoteId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
