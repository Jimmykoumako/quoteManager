package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB // Package-level variable to hold the Gorm DB instance

const AdminRole = "admin"

// // SetDB sets the Gorm DB instance for the user model
// func SetDB(database *gorm.DB) {
//     db = database
// }

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

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null;index"` // Added indexing for better query performance
	Password  string    `json:"-" gorm:"not null"`
	Quotes []Quote `json:"quotes" gorm:"foreignKey:user_id"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// BeforeCreate hooks into the GORM lifecycle to hash the user's password before saving
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies if the provided password matches the hashed password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// RegisterUser registers a new user in the database
func RegisterUser(input UserInput) (User, error) {
	newUser := User{
		Username: input.Username,
		Password: input.Password,
	}

	// Save newUser to the database (using Gorm or your preferred ORM)
	result := db.Create(&newUser)
    if result.Error != nil {
        // Handle the error, for example:
        return User{}, result.Error
    }

	return newUser, nil
}

// AuthenticateUser authenticates a user based on login credentials
func AuthenticateUser(input LoginInput) (User, error) {
	// Retrieve user from the database by username
	foundUser, err := GetUserByUsername(input.Username)
	if err != nil {
		// Handle the case where the user is not found or other database-related errors
		return User{}, err
	}

	// Check if the provided password matches the hashed password
	if err := foundUser.CheckPassword(input.Password); err != nil {
		return User{}, err
	}

	return foundUser, nil
}

// GetUserByUsername retrieves a user from the database by username
func GetUserByUsername(username string) (User, error) {
	var user User

	// Replace "db" with your Gorm DB instance
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		// Handle specific cases, such as gorm.ErrRecordNotFound
		return User{}, result.Error
	}

	return user, nil
}

// UpdateUser updates user details in the database
func UpdateUser(userID string, updatedUser User) (User, error) {
    // Replace "db" with your actual Gorm DB instance
    result := db.Model(&User{}).Where("id = ?", userID).Updates(updatedUser)
    if result.Error != nil {
        return User{}, result.Error
    }

    return updatedUser, nil
}

// DeleteUser deletes a user account and associated data
func DeleteUser(userID string) error {
    // Replace "db" with your actual Gorm DB instance
    result := db.Where("id = ?", userID).Delete(&User{})
    if result.Error != nil {
        return result.Error
    }

    // Delete associated quotes
    result = db.Where("user_id = ?", userID).Delete(&Quote{})
    if result.Error != nil {
        return result.Error
    }

    // Delete associated folders
    result = db.Where("user_id = ?", userID).Delete(&Folder{})
    if result.Error != nil {
        return result.Error
    }

    return nil
}


// GetUserByID retrieves a user by ID with associated quotes and folders
func GetUserByID(userID string) (User, error) {
    var user User
    result := db.Preload("Quotes").Preload("Folders").First(&user, userID)
    if result.Error != nil {
        return User{}, result.Error
    }
    return user, nil
}

// GetQuotesByUserID retrieves quotes for a specific user by ID
func GetQuotesByUserID(userID string) ([]Quote, error) {
    // Implement logic to fetch quotes for the specified user from the database
    // Example using Gorm
    var quotes []Quote
    result := db.Where("user_id = ?", userID).Find(&quotes)
    if result.Error != nil {
        return nil, result.Error
    }
    return quotes, nil
}

// GetFoldersByUserID retrieves folders for a specific user by ID
func GetFoldersByUserID(userID string) ([]Folder, error) {
    // Implement logic to fetch folders for the specified user from the database
    // Example using Gorm
    var folders []Folder
    result := db.Where("user_id = ?", userID).Find(&folders)
    if result.Error != nil {
        return nil, result.Error
    }
    return folders, nil
}


