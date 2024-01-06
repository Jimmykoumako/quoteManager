// database/literary_work.go
package database

import (
	"errors"
	"api/models"
	"api/logger"
)

// LiteraryWorkInput represents the input for creating or updating a literary work
type LiteraryWorkInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author"`
}

// CreateLiteraryWork creates a new literary work in the database
func CreateLiteraryWork(input LiteraryWorkInput) (models.LiteraryWork, error) {
	logger.Log.Info("Welcome to database.CreateLiteraryWork")

	if db == nil {
		return models.LiteraryWork{}, errors.New("nil database provided")
	}

	// Implement your logic to create a new literary work

	newLiteraryWork := models.LiteraryWork{
		Title:  input.Title,
		Author: input.Author,
	}

	// Save newLiteraryWork to the database (using Gorm or your preferred ORM)
	result := db.Create(&newLiteraryWork)
	if result.Error != nil {
		// Handle the error, for example:
		return models.LiteraryWork{}, result.Error
	}

	logger.Log.Info("Bye from database.CreateLiteraryWork")
	return newLiteraryWork, nil
}

// GetLiteraryWorks returns a list of all literary works from the database
func GetLiteraryWorks() ([]models.LiteraryWork, error) {
	logger.Log.Info("Welcome to database.GetLiteraryWorks")

	if db == nil {
		return nil, errors.New("nil database provided")
	}

	// Implement your logic to fetch all literary works

	var literaryWorks []models.LiteraryWork
	result := db.Find(&literaryWorks)
	if result.Error != nil {
		// Handle the error, for example:
		return nil, result.Error
	}

	logger.Log.Info("Bye from database.GetLiteraryWorks")
	return literaryWorks, nil
}

// GetLiteraryWorkByID returns a literary work with the specified ID from the database
func GetLiteraryWorkByID(id uint) (models.LiteraryWork, error) {
	logger.Log.Info("Welcome to database.GetLiteraryWorkByID")

	if db == nil {
		return models.LiteraryWork{}, errors.New("nil database provided")
	}

	// Implement your logic to fetch a literary work by ID

	var literaryWork models.LiteraryWork
	result := db.First(&literaryWork, id)
	if result.Error != nil {
		// Handle the error, for example:
		return models.LiteraryWork{}, result.Error
	}

	logger.Log.Info("Bye from database.GetLiteraryWorkByID")
	return literaryWork, nil
}

// UpdateLiteraryWork updates an existing literary work in the database
func UpdateLiteraryWork(id uint, input LiteraryWorkInput) (models.LiteraryWork, error) {
	logger.Log.Info("Welcome to database.UpdateLiteraryWork")

	if db == nil {
		return models.LiteraryWork{}, errors.New("nil database provided")
	}

	// Implement your logic to update a literary work

	var existingLiteraryWork models.LiteraryWork
	result := db.First(&existingLiteraryWork, id)
	if result.Error != nil {
		// Handle the error, for example:
		return models.LiteraryWork{}, result.Error
	}

	// Update the literary work details
	existingLiteraryWork.Title = input.Title
	existingLiteraryWork.Author = input.Author

	// Save the updated literary work to the database
	result = db.Save(&existingLiteraryWork)
	if result.Error != nil {
		// Handle the error, for example:
		return models.LiteraryWork{}, result.Error
	}

	logger.Log.Info("Bye from database.UpdateLiteraryWork")
	return existingLiteraryWork, nil
}

// DeleteLiteraryWork deletes a literary work with the specified ID from the database
func DeleteLiteraryWork(id uint) error {
	logger.Log.Info("Welcome to database.DeleteLiteraryWork")

	if db == nil {
		return errors.New("nil database provided")
	}

	// Implement your logic to delete a literary work

	var existingLiteraryWork models.LiteraryWork
	result := db.First(&existingLiteraryWork, id)
	if result.Error != nil {
		// Handle the error, for example:
		return result.Error
	}

	// Delete the literary work from the database
	result = db.Delete(&existingLiteraryWork)
	if result.Error != nil {
		// Handle the error, for example:
		return result.Error
	}

	logger.Log.Info("Bye from database.DeleteLiteraryWork")
	return nil
}
