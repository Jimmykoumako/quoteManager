package controllers

import (
	"api/database"
	"api/models"
	"api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

// RefreshToken refreshes the access token using a refresh token
func RefreshToken(c *gin.Context) {
	// Extract refresh token from request
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	// Validate refresh token (check if it's valid and not expired)
	valid, userID, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// If valid, generate a new access token
	newAccessToken, err := utils.GenerateJWTwithID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	// Set the new access token as a cookie
	c.SetCookie("jwt", newAccessToken, int(utils.AccessTokenExpiration.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

// AuthenticateUser authenticates a user based on login credentials
func AuthenticateUser(input database.LoginInput) (models.User, error) {
	// Retrieve user from the database by username
	foundUser, err := database.GetUserByUsername(input.Username)
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

// CheckUsernameAvailability checks if a username is available
func CheckUsernameAvailability(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the username is available
    available, err := database.UsernameExists(input.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"available": !available})
}

// RegisterBasicUserInfo collects basic user information (username and password)
func RegisterBasicUserInfo(c *gin.Context) {
    var userInput struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
        return
    }

    // Check if the username is available
    available, _ := database.UsernameExists(userInput.Username)
    if !available {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already in use"})
        return
    }

    if err := database.SaveBasicUserInfoToDatabase(userInput.Username, userInput.Password); err != nil {
		log.Println("Failed to register user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}
	
    // Return success response
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Basic information registered successfully"})
}

// RegisterAdditionalUserInfo collects additional user information during registration
func RegisterAdditionalUserInfo(c *gin.Context) {
    var userInput  database.UserRegistrationInput

    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    // Save additional user information to the database (replace with your database logic)
    if err := database.SaveAdditionalUserInfoToDatabase(userInput.Username, userInput.Email, userInput.FirstName, userInput.LastName, userInput.Birthdate); err != nil {
        log.Println("Failed to register user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
        return
    }

    // Return success response
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Additional information registered successfully"})
}

// RegisterFinalize completes the user registration
func RegisterFinalize(c *gin.Context) {
    var input database.UserRegistrationInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
    // Save the combined user information
    newUser, newProfile, err := database.RegisterUserAndProfile(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete registration"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Registration completed successfully", "user": newUser, "profile": newProfile})
}

// LoginUser authenticates a user and generates a JWT
func LoginUser(c *gin.Context) {
	var loginInput database.LoginInput

	// Bind the JSON request body to the LoginInput struct
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement logic to authenticate the user based on the login credentials
	authUser, err := database.AuthenticateUser(loginInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(authUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the JWT as a cookie with a one-hour expiration
	c.SetCookie("jwt", token, 3600, "/", "", false, true)

	// Respond with other data or a success message
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
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
	result, err := database.UpdateUser(userID, updatedUser)
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

	err := database.DeleteUser(userID)
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

	quotes, err := database.GetQuotesByUserID(userID)
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
	folders, err := database.GetFoldersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, folders)
}

// GetQuotesByUserID retrieves quotes for a specific user
func GetQuotesByUserID(c *gin.Context) {
	userID := c.Param("id")

	quotes, err := database.GetQuotesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quotes)
}

// GetFoldersByUserID retrieves folders for a specific user
func GetFoldersByUserID(c *gin.Context) {
	userID := c.Param("id")

	folders, err := database.GetFoldersByUserID(userID)
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
	result, err := database.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUserProfile retrieves the profile information of a specific user
func GetUserProfile(c *gin.Context) {
    userID := c.Param("id")

    // Implement logic to fetch user profile information from the database
    userProfile, err := database.GetUserProfile(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
        return
    }

    c.JSON(http.StatusOK, userProfile)
}






