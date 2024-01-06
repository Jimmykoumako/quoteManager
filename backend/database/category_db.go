// database/category.go
package database

import (
	"errors"
	"api/models"
	"log"
	"api/logger"
)


// CategoryInput represents the input for creating a new category
type CategoryInput struct {
	Name string `json:"name"`
}

// CreateCategory creates a new category in the database
func CreateCategory(input CategoryInput) (models.Category, error) {
	logger.Log.Info("Welcome to database.CreateCategory")

	if db == nil {
		return models.Category{}, errors.New("nil database provided")
	}

	// Implement your logic to create a new category

	newCategory := models.Category{
		Name: input.Name,
	}

	// Save newCategory to the database (using Gorm or your preferred ORM)
	result := db.Create(&newCategory)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error creating category: %v", result.Error)
		return models.Category{}, result.Error
	}

	logger.Log.Info("Bye from models.CreateCategory")
	return newCategory, nil
}

// GetCategories returns a list of all categories from the database
func GetCategories() ([]models.Category, error) {
	logger.Log.Info("Welcome to database.GetCategories")

	if db == nil {
		return nil, errors.New("nil database provided")
	}

	// Implement your logic to fetch all categories

	var categories []models.Category
	result := db.Find(&categories)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error while getting categories: %v", result.Error)
		return nil, result.Error
	}

	logger.Log.Info("Bye from models.GetCategories")
	return categories, nil
}

// GetCategoryByID returns a category with the specified ID from the database
func GetCategoryByID(id uint) (models.Category, error) {
	logger.Log.Info("Welcome to database.GetCategoryByID")

	if db == nil {
		return models.Category{}, errors.New("nil database provided")
	}

	// Implement your logic to fetch a category by ID

	var category models.Category
	result := db.First(&category, id)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error getting category: %v", result.Error)
		return models.Category{}, result.Error
	}

	logger.Log.Info("Bye from models.GetCategoryByID")
	return category, nil
}

// UpdateCategory updates an existing category in the database
func UpdateCategory(id uint, input CategoryInput) (models.Category, error) {
	logger.Log.Info("Welcome to database.UpdateCategory")

	if db == nil {
		return models.Category{}, errors.New("nil database provided")
	}

	// Implement your logic to update a category

	var existingCategory models.Category
	result := db.Where("name = ?", input.Name).First(&existingCategory)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error getting category: for update %v", result.Error)
		return models.Category{}, result.Error
	}

	// Update the category name
	existingCategory.Name = input.Name

	// Save the updated category to the database
	result = db.Save(&existingCategory)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error updating category: %v", result.Error)
		return models.Category{}, result.Error
	}

	logger.Log.Info("Bye from database.UpdateCategory")
	return existingCategory, nil
}

// DeleteCategory deletes a category with the specified ID from the database
func DeleteCategory(id uint) error {
	logger.Log.Info("Welcome to database.DeleteCategory")

	if db == nil {
		return errors.New("nil database provided")
	}

	// Implement your logic to delete a category

	var existingCategory models.Category
	result := db.First(&existingCategory, id)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error getting category to delete: %v", result.Error)
		return result.Error
	}

	// Delete the category from the database
	result = db.Delete(&existingCategory)
	if result.Error != nil {
		// Handle the error, for example:
		log.Printf("Error deleting category: %v", result.Error)
		return result.Error
	}

	logger.Log.Info("Bye from models.DeleteCategory")
	return nil
}
