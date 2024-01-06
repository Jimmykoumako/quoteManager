package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"
	"api/database"
)



// CreateCategory creates a new category in the database
func CreateCategory(c *gin.Context) {
	fmt.Println("Welcome to controllers.CreateCategory")

	var input database.CategoryInput

	// Bind the JSON request body to the CategoryInput struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Example of basic input validation
	if len(input.Name) == 0 {
		// Handle invalid input, e.g., return an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty"})
		return
	}

	// Call the model function to create a new category
	newCategory, err := database.CreateCategory(input)
	if err != nil {
		if strings.Contains(err.Error(), "category name already in use") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category name already in use"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

// GetCategories returns a list of all categories from the database
func GetCategories(c *gin.Context) {
	fmt.Println("Welcome to controllers.GetCategories")

	// Call the model function to get all categories
	categories, err := database.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID returns a category with the specified ID from the database
func GetCategoryByID(c *gin.Context) {
	fmt.Println("Welcome to controllers.GetCategoryByID")

	categoryID := c.Param("id")

	// Convert the categoryID to uint
	id, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Call the model function to get the category by ID
	category, err := database.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates an existing category in the database
func UpdateCategory(c *gin.Context) {
	fmt.Println("Welcome to controllers.UpdateCategory")

	categoryID := c.Param("id")

	// Convert the categoryID to uint
	id, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var input database.CategoryInput

	// Bind the JSON request body to the CategoryInput struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Call the model function to update the category
	updatedCategory, err := database.UpdateCategory(uint(id), input)
	if err != nil {
		if strings.Contains(err.Error(), "category not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory deletes a category with the specified ID from the database
func DeleteCategory(c *gin.Context) {
	fmt.Println("Welcome to controllers.DeleteCategory")

	categoryID := c.Param("id")

	// Convert the categoryID to uint
	id, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Call the model function to delete the category
	err = database.DeleteCategory(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "category not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
