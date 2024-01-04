package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm" 
	"net/http"
	"api/models"
	"api/utils"
)

var DB *gorm.DB

// GetFoldersForUser returns all folders for the authenticated user
func GetFoldersForUser(c *gin.Context) {
	// For demonstration purposes, let's assume authentication is done and the user ID is available
	userID := utils.GetUserIDFromContext(c) // Replace with actual user ID retrieval

	// Implement logic to fetch folders for the authenticated user from the database
	var folders []models.Folder
	if err := models.DB.Where("UserID = ?", userID).Find(&folders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, folders)
}

// GetFolderByID returns a specific folder by ID
func GetFolderByID(c *gin.Context) {
	// Extract folder ID from the URL parameter
	folderID := c.Param("id")

	// Implement logic to fetch a folder by ID from the database
	var folder models.Folder
	if err := models.DB.First(&folder, folderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}

	c.JSON(http.StatusOK, folder)
}

/// CreateFolder creates a new folder
func CreateFolder(c *gin.Context) {
	// For demonstration purposes, let's assume authentication is done and the user ID is available
	userID := utils.GetUserIDFromContext(c)

	// Convert userID to uint if needed
	userIDUint, err := utils.ConvertUserIDToUint(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Implement logic to create a new folder
	var folder models.Folder
	if err := c.ShouldBindJSON(&folder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the user ID for the folder
	folder.UserID = userIDUint

	// Example using Gorm
	if err := models.DB.Create(&folder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, folder)
}

// UpdateFolder updates an existing folder by ID
func UpdateFolder(c *gin.Context) {
	// Extract folder ID from the URL parameter
	folderID := c.Param("id")

	// Implement logic to update a folder by ID in the database
	var updatedFolder models.Folder
	if err := c.ShouldBindJSON(&updatedFolder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example using Gorm
	if err := models.DB.Model(&models.Folder{}).Where("id = ? AND UserID = ?", folderID, utils.GetUserIDFromContext(c)).Updates(&updatedFolder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedFolder)
}

// DeleteFolder deletes a folder by ID
func DeleteFolder(c *gin.Context) {
	// Extract folder ID from the URL parameter
	folderID := c.Param("id")

	// Implement logic to delete a folder by ID from the database
	// Example using Gorm
	if err := models.DB.Where("id = ? AND UserID = ?", folderID, utils.GetUserIDFromContext(c)).Delete(&models.Folder{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Folder deleted successfully"})
}
