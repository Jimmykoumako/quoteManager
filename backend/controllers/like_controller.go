package controllers

import (
	"api/database"
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLikeByID returns a specific like by ID
func GetLikeByID(c *gin.Context) {
	likeID := c.Param("id")

	// Implement logic to fetch a like by ID from the database
	like, err := database.GetLikeByID(likeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, like)
}

// AddLike adds a new like
func AddLike(c *gin.Context) {
	var like models.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement logic to add a new like to the database
	err := database.AddLike(&like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, like)
}

// UpdateLike updates an existing like by ID
func UpdateLike(c *gin.Context) {
	likeID := c.Param("id")

	var updatedLike models.Like
	if err := c.ShouldBindJSON(&updatedLike); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement logic to update a like by ID in the database
	err := database.UpdateLike(likeID, &updatedLike)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLike)
}

// DeleteLike deletes a like by ID
func DeleteLike(c *gin.Context) {
	likeID := c.Param("id")

	// Implement logic to delete a like by ID from the database
	err := database.DeleteLike(likeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}
