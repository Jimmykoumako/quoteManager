// database/tag.go
package database

import (
	"fmt"
	"errors"
	"api/models"
)

// TagInput represents the input for creating or updating a tag
type TagInput struct {
	Name string `json:"name"`
}

// GetTags returns a list of all tags from the database
func GetTags() ([]models.Tag, error) {
	fmt.Println("Welcome to database.GetTags")

	if db == nil {
		return nil, errors.New("nil database provided")
	}

	var tags []models.Tag
	result := db.Find(&tags)
	if result.Error != nil {
		// Handle the error, for example:
		return nil, result.Error
	}

	fmt.Println("Bye from database.GetTags")
	return tags, nil
}

// GetTagByID returns a tag with the specified ID from the database
func GetTagByID(id uint) (models.Tag, error) {
	fmt.Println("Welcome to database.GetTagByID")

	if db == nil {
		return models.Tag{}, errors.New("nil database provided")
	}

	var tag models.Tag
	result := db.First(&tag, id)
	if result.Error != nil {
		// Handle the error, for example:
		return models.Tag{}, result.Error
	}

	fmt.Println("Bye from database.GetTagByID")
	return tag, nil
}

// CreateTag creates a new tag in the database
func CreateTag(input TagInput) (models.Tag, error) {
	fmt.Println("Welcome to database.CreateTag")

	if db == nil {
		return models.Tag{}, errors.New("nil database provided")
	}

	// Check if the tag name already exists in the database
	var existingTag models.Tag
	if err := db.Where("name = ?", input.Name).First(&existingTag).Error; err == nil {
		// Tag name already in use, return an error
		return models.Tag{}, errors.New("tag name already in use")
	}

	newTag := models.Tag{
		Name: input.Name,
	}

	// Save newTag to the database (using Gorm or your preferred ORM)
	result := db.Create(&newTag)
	if result.Error != nil {
		// Handle the error, for example:
		return models.Tag{}, result.Error
	}

	fmt.Println("Bye from database.CreateTag")
	return newTag, nil
}

// UpdateTag updates an existing tag in the database
func UpdateTag(id uint, input TagInput) (models.Tag, error) {
	fmt.Println("Welcome to database.UpdateTag")

	if db == nil {
		return models.Tag{}, errors.New("nil database provided")
	}

	// Check if the tag with the specified ID exists
	var existingTag models.Tag
	result := db.First(&existingTag, id)
	if result.Error != nil {
		// Handle the error, for example:
		return models.Tag{}, errors.New("tag not found")
	}

	// Update the tag name
	existingTag.Name = input.Name

	// Save the updated tag to the database
	result = db.Save(&existingTag)
	if result.Error != nil {
		// Handle the error, for example:
		return models.Tag{}, result.Error
	}

	fmt.Println("Bye from database.UpdateTag")
	return existingTag, nil
}

// DeleteTag deletes a tag with the specified ID from the database
func DeleteTag(id uint) error {
	fmt.Println("Welcome to database.DeleteTag")

	if db == nil {
		return errors.New("nil database provided")
	}

	// Check if the tag with the specified ID exists
	var existingTag models.Tag
	result := db.First(&existingTag, id)
	if result.Error != nil {
		// Handle the error, for example:
		return errors.New("tag not found")
	}

	// Delete the tag from the database
	result = db.Delete(&existingTag)
	if result.Error != nil {
		// Handle the error, for example:
		return result.Error
	}

	fmt.Println("Bye from database.DeleteTag")
	return nil
}
