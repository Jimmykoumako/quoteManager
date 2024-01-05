package database

import (
	"gorm.io/gorm"
	"api/models"
)



// SetDB sets the Gorm database connection
func SetDB(database *gorm.DB) {
	db = database
}

// GetQuotes fetches all quotes from the database
func GetQuotes() ([]models.Quote, error) {
	var quotes []models.Quote
	if err := db.Find(&quotes).Error; err != nil {
		return nil, err
	}
	return quotes, nil
}

// GetQuoteByID fetches a specific quote by ID from the database
func GetQuoteByID(quoteID uint) (models.Quote, error) {
	var quote models.Quote
	if err := db.First(&quote, quoteID).Error; err != nil {
		return models.Quote{}, err
	}
	return quote, nil
}

// AddQuote adds a new quote to the database
// AddQuote adds a new quote to the database
func AddQuote(newQuote models.Quote) (models.Quote, error) {
	// Validate required fields
	if newQuote.Text == "" || newQuote.Author == "" {
		return models.Quote{}, ErrInvalidPayload
	}

	// Set default values for optional fields
	if newQuote.Category == "" {
		newQuote.Category = "Default Category"
	}

	// Set default value for Tags if not provided
	if len(newQuote.Tags) == 0 {
		newQuote.Tags = []string{"Default Tag"}
	}

	// Set default values for optional fields
	if newQuote.WorkID  != 1 {
		newQuote.WorkID  = 1
	}

	// Add the new quote to the database
	if err := db.Create(&newQuote).Error; err != nil {
		return models.Quote{}, err
	}

	return newQuote, nil
}


// UpdateQuote updates an existing quote by ID in the database
func UpdateQuote(quoteID uint, updatedQuote models.Quote) (models.Quote, error) {
	if err := db.First(&models.Quote{}, quoteID).Updates(updatedQuote).Error; err != nil {
		return models.Quote{}, err
	}
	return updatedQuote, nil
}

// DeleteQuote deletes a quote by ID from the database
func DeleteQuote(quoteID uint) error {
	if err := db.Delete(&models.Quote{}, quoteID).Error; err != nil {
		return err
	}
	return nil
}
