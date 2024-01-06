// controllers/user_profile_controller.go

package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "api/models"
    "api/database"
)

// GetUserProfileByID retrieves a user profile by ID
func GetUserProfileByID(c *gin.Context) {
    profileID := c.Param("id")

    // Convert profileID to uint
    profileIDUint, err := convertToUint(profileID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
        return
    }

    // Fetch user profile from the database
    userProfile, err := database.GetUserProfileByID(profileIDUint)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
        return
    }

    c.JSON(http.StatusOK, userProfile)
}

// CreateUserProfile creates a new user profile
func CreateUserProfile(c *gin.Context) {
    var newProfile models.UserProfile

    // Bind JSON request body to UserProfile model
    if err := c.ShouldBindJSON(&newProfile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Add the new user profile to the database
    if err := database.CreateUserProfile(&newProfile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
        return
    }

    c.JSON(http.StatusCreated, newProfile)
}

// UpdateUserProfile updates an existing user profile by ID
func UpdateUserProfile(c *gin.Context) {
    profileID := c.Param("id")

    // Convert profileID to uint
    profileIDUint, err := convertToUint(profileID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
        return
    }

    var updatedProfile models.UserProfile

    // Bind JSON request body to updated UserProfile model
    if err := c.ShouldBindJSON(&updatedProfile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Update the user profile in the database
    if err := database.UpdateUserProfile(profileIDUint, &updatedProfile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
        return
    }

    c.JSON(http.StatusOK, updatedProfile)
}

// DeleteUserProfile deletes a user profile by ID
func DeleteUserProfile(c *gin.Context) {
    profileID := c.Param("id")

	// Convert profileID to uint
	profileIDUint, err := convertToUint(profileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	// Delete the user profile from the database
	if err := database.DeleteUserProfile(profileIDUint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile deleted successfully"})
}
	
