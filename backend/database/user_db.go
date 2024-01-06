package database

import (
	"api/models"
	"errors"
	"fmt"
)

// UserInput represents the input data for user registration
type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginInput represents the input data for user login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUser registers a new user in the database
func RegisterUser(input UserInput) (models.User, error) {
	fmt.Println("Welcome to models.RegisterUser")
	if db == nil {
		return models.User{}, errors.New("nil database provided")
	}

	// Check if the username already exists in the database
	var existingUser models.User
	if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		// Username already exists, return an error
		return models.User{}, errors.New("username already in use")
	}

	newUser := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	// Save newUser to the database (using Gorm or your preferred ORM)
	result := db.Create(&newUser)
	if result.Error != nil {
		// Handle the error, for example:
		return models.User{}, result.Error
	}

	fmt.Println("Bye from models.RegisterUser")
	return newUser, nil
}

// AuthenticateUser authenticates a user based on login credentials
func AuthenticateUser(input LoginInput) (models.User, error) {
	// Retrieve user from the database by username
	foundUser, err := GetUserByUsername(input.Username)
	if err != nil {
		// Handle the case where the user is not found or other database-related errors
		return models.User{}, err
	}

	// Check if the provided password matches the hashed password
	if err := foundUser.CheckPassword(input.Password); err != nil {
		return models.User{}, err
	}

	return foundUser, nil
}

// GetUserByUsername retrieves a user from the database by username
func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	// Replace "db" with your Gorm db instance
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		// Handle specific cases, such as gorm.ErrRecordNotFound
		return models.User{}, result.Error
	}

	return user, nil
}

// UpdateUser updates user details in the database
func UpdateUser(userID string, updatedUser models.User) (models.User, error) {
	// Replace "db" with your actual Gorm db instance
	result := db.Model(&models.User{}).Where("id = ?", userID).Updates(updatedUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return updatedUser, nil
}

// DeleteUser deletes a user account and associated data
func DeleteUser(userID string) error {
	// Replace "db" with your actual Gorm db instance
	result := db.Where("id = ?", userID).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}

	// Delete associated quotes
	result = db.Where("user_id = ?", userID).Delete(&models.Quote{})
	if result.Error != nil {
		return result.Error
	}

	// Delete associated folders
	result = db.Where("user_id = ?", userID).Delete(&models.Folder{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetUserByID retrieves a user by ID with associated quotes and folders
func GetUserByID(userID string) (models.User, error) {
	var user models.User
	result := db.Preload("Quotes").Preload("Folders").First(&user, userID)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

// GetQuotesByUserID retrieves quotes for a specific user by ID
func GetQuotesByUserID(userID string) ([]models.Quote, error) {
	// Implement logic to fetch quotes for the specified user from the database
	// Example using Gorm
	var quotes []models.Quote
	result := db.Where("user_id = ?", userID).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// GetFoldersByUserID retrieves folders for a specific user by ID
func GetFoldersByUserID(userID string) ([]models.Folder, error) {
	// Implement logic to fetch folders for the specified user from the database
	// Example using Gorm
	var folders []models.Folder
	result := db.Where("user_id = ?", userID).Find(&folders)
	if result.Error != nil {
		return nil, result.Error
	}
	return folders, nil
}
