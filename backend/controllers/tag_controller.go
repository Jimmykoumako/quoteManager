// controllers/tag_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"api/logger"
	"github.com/gin-gonic/gin"
	"api/database"
)


// GetTags returns a list of all tags from the database
func GetTags(c *gin.Context) {
	logger.Log.Info("Welcome to controllers.GetTags")

	tags, err := database.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetTagByID returns a tag with the specified ID from the database
func GetTagByID(c *gin.Context) {
	logger.Log.Info("Welcome to controllers.GetTagByID")

	tagID := c.Param("id")

	id, err := strconv.ParseUint(tagID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	tag, err := database.GetTagByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// CreateTag creates a new tag in the database
func CreateTag(c *gin.Context) {
	logger.Log.Info("Welcome to controllers.CreateTag")

	var input database.TagInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	newTag, err := database.CreateTag(input)
	if err != nil {
		if strings.Contains(err.Error(), "tag name already in use") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag name already in use"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, newTag)
}

// UpdateTag updates an existing tag in the database
func UpdateTag(c *gin.Context) {
	logger.Log.Info("Welcome to controllers.UpdateTag")

	tagID := c.Param("id")

	id, err := strconv.ParseUint(tagID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	var input database.TagInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	updatedTag, err := database.UpdateTag(uint(id), input)
	if err != nil {
		if strings.Contains(err.Error(), "tag not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update tag"})
		return
	}

	c.JSON(http.StatusOK, updatedTag)
}

// DeleteTag deletes a tag with the specified ID from the database
func DeleteTag(c *gin.Context) {
	logger.Log.Info("Welcome to controllers.DeleteTag")

	tagID := c.Param("id")

	id, err := strconv.ParseUint(tagID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	err = database.DeleteTag(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tag not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
