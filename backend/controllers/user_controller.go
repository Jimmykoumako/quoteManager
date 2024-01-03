package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"api/models"
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