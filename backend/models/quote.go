package models

import (
	"time"
)

// Quote Entity
type Quote struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Text      string    `json:"text" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags" gorm:"-"`
	Feedback  []Feedback `json:"feedback" gorm:"foreignKey:QuoteID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// GetQuotes returns the quotes associated with a Quote entity
func (q Quote) GetQuotes() []Quote {
	return []Quote{q}
}

