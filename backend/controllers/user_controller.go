package controllers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

// AuthenticateUser authenticates a user based on login credentials
func AuthenticateUser(input models.LoginInput) (models.User, error) {
	// Retrieve user from the database by username
	foundUser, err := models.GetUserByUsername(input.Username)
	if err != nil {
		// Handle the case where the user is not found
		return models.User{}, err
	}

	// Check if the provided password matches the hashed password
	if err := foundUser.CheckPassword(input.Password); err != nil {
		// Handle the case where the password is incorrect
		return models.User{}, err
	}

	return foundUser, nil
}

// RegisterUser registers a new user
func RegisterUser(c *gin.Context) {
	var userInput models.UserInput

	// Bind the JSON request body to the UserInput struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Implement logic to validate and register the new user in the database
	newUser, err := models.RegisterUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}


// LoginUser authenticates a user
func LoginUser(c *gin.Context) {
	var loginInput models.LoginInput

	// Bind the JSON request body to the LoginInput struct
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement logic to authenticate the user based on the login credentials
	authUser, err := models.AuthenticateUser(loginInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, authUser)
}


// UpdateUser updates an existing user by ID
func UpdateUser(c *gin.Context) {
    // Extract user ID from the URL parameter
    userID := c.Param("id")

    // Parse the authenticated user ID from the context or token
    authUserID, err := parseUserIDFromContext(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Ensure that the authenticated user is updating their own details
    if userID != authUserID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
        return
    }

    // Implement logic to update a user by ID in the database
    var updatedUser models.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Example using Gorm
    result, err := models.UpdateUser(userID, updatedUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    c.JSON(http.StatusOK, result)
}

// parseUserIDFromContext extracts the user ID from the request context
func parseUserIDFromContext(c *gin.Context) (string, error) {
    userIDValue, exists := c.Get("userID")
    if !exists {
        return "", fmt.Errorf("User ID not found in request context")
    }

    userID, ok := userIDValue.(string)
    if !ok {
        return "", fmt.Errorf("User ID is not of type string")
    }

    return userID, nil
}







// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
    // Extract user ID from the URL parameter
    userID := c.Param("id")


    err := models.DeleteUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}



// GetUserQuotes retrieves quotes for a specific user
func GetUserQuotes(c *gin.Context) {
	// Extract user ID from the URL parameter
	userID := c.Param("id")

	// Implement logic to fetch quotes for the specified user from the database
	// Example using Gorm
	quotes, err := models.GetQuotesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quotes)
}

// GetUserFolders retrieves folders for a specific user
func GetUserFolders(c *gin.Context) {
	// Extract user ID from the URL parameter
	userID := c.Param("id")

	// Implement logic to fetch folders for the specified user from the database
	// Example using Gorm
	folders, err := models.GetFoldersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, folders)
}


// GetQuotesByUserID retrieves quotes for a specific user
func GetQuotesByUserID(c *gin.Context) {
    userID := c.Param("id")

    quotes, err := models.GetQuotesByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quotes)
}

// GetFoldersByUserID retrieves folders for a specific user
func GetFoldersByUserID(c *gin.Context) {
    userID := c.Param("id")

    folders, err := models.GetFoldersByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, folders)
}


// GetUserByID retrieves a user by ID with associated quotes and folders
func GetUserByID(c *gin.Context) {
    userID := c.Param("id")

    //var user models.User
    result, err := models.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, result)
}

