// database.go
package database

import (
	"api/models"
)

// SearchAndFilterQuotes fetches quotes based on search and filters
func SearchAndFilterQuotes(searchQuery, authorFilter, categoryFilter string) ([]models.Quote, error) {
	var quotes []models.Quote

	// Start building the query
	query := db.Model(&models.Quote{})

	// Apply search filter
	if searchQuery != "" {
		query = query.Where("title LIKE ?", "%"+searchQuery+"%")
	}

	// Apply author filter
	if authorFilter != "" {
		query = query.Where("author = ?", authorFilter)
	}

	// Apply category filter
	if categoryFilter != "" {
		query = query.Joins("JOIN categories ON quotes.category_id = categories.id").
			Where("categories.name = ?", categoryFilter)
	}

	// Execute the query
	if err := query.Find(&quotes).Error; err != nil {
		return nil, err
	}

	return quotes, nil
}
