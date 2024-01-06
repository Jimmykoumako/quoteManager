package database

import (
	"golang.org/x/crypto/bcrypt"
	"api/models"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)


// UserInput represents the input data for user registration
type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserRegistrationInput holds the data collected during user registration
type UserRegistrationInput struct {
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	// Add other fields as needed
}

// LoginInput represents the input data for user login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UsernameExists checks if a username already exists in the database
func UsernameExists(username string) (bool, error) {
    var user models.User
    result := db.Where("username = ?", username).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            // User not found, username is available
            return false, nil
        }
        // An error occurred while querying the database
        return false, result.Error
    }

    // User found, username is not available
    return true, nil
}


// SaveBasicUserInfoToDatabase saves basic user information to the database
func SaveBasicUserInfoToDatabase(username, password string) error {
    // Hash the password before saving it to the database
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password: %v", err)
        return err
    }

    // Create input for RegisterUserAndProfile
    input := UserRegistrationInput{
        Username:  username,
        Password:  string(hashedPassword),
        FirstName: "", // You can set default values or leave them empty
        LastName:  "",
        Email:     "",
        Birthdate: time.Time{},
    }

    // Call RegisterUserAndProfile to create user and profile
    _, _, err = RegisterUserAndProfile(input)
    if err != nil {
        log.Printf("Error registering user and profile: %v", err)
    }

    return err
}

// SaveAdditionalUserInfoToDatabase saves additional user information to the database
func SaveAdditionalUserInfoToDatabase(username, email, firstName, lastName string, birthdate time.Time) error {
    // Create input for RegisterUserAndProfile
    input := UserRegistrationInput{
        Username:  username,
        Password:  "", // You can set default values or leave them empty
        FirstName: firstName,
        LastName:  lastName,
        Email:     email,
        Birthdate: birthdate,
    }

    // Call RegisterUserAndProfile to create user and profile
    _, _, err := RegisterUserAndProfile(input)
    if err != nil {
        log.Printf("Error registering user and profile: %v", err)
    }

    return err
}


// RegisterUserAndProfile registers a new user and creates a user profile in the database
func RegisterUserAndProfile(input UserRegistrationInput) (models.User, models.UserProfile, error) {
    // Check if the username already exists in the database
    _, err := GetUserByUsername(input.Username)
    if err == nil {
        // Username already exists, return an error
        return models.User{}, models.UserProfile{}, errors.New("username already in use")
    }

    newUser := models.User{
        Username: input.Username,
        Password: input.Password,
    }

    // Save newUser to the database
    result := db.Create(&newUser)
    if result.Error != nil {
        return models.User{}, models.UserProfile{}, result.Error
    }

    // Create a user profile for the new user
    newProfile := models.UserProfile{
        UserID:    newUser.ID,
        FirstName: input.FirstName,
        LastName:  input.LastName,
        Email:     input.Email,
        Birthdate: input.Birthdate,
    }

    // Save the new profile to the database
    result = db.Create(&newProfile)
    if result.Error != nil {
        return models.User{}, models.UserProfile{}, result.Error
    }

    return newUser, newProfile, nil
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

// GetUserProfile fetches the profile information of a specific user
func GetUserProfile(userID string) (models.UserProfile, error) {
    // Implement logic to fetch user profile information from the database
    // Example using Gorm
    var userProfile models.UserProfile
    if err := db.Where("UserID = ?", userID).First(&userProfile).Error; err != nil {
        return models.UserProfile{}, err
    }

    return userProfile, nil
}
