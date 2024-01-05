package database

import (
	"api/models"
	"fmt"
	"time"
)

func createQuotesForUser(userID uint) []models.Quote {
	// Customize the number of quotes per user
	numQuotes := 5
	var quotes []models.Quote

	for i := 1; i <= numQuotes; i++ {
		quote := models.Quote{
			Text:      fmt.Sprintf("Quote %d for User %d", i, userID),
			Author:    "Anonymous",
			UserID:    userID,
			Category:  "General",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		quotes = append(quotes, quote)
	}

	return quotes
}
