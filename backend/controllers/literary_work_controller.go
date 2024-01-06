// controllers/literary_work_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api/database"
)



// CreateLiteraryWork creates a new literary work in the database
func CreateLiteraryWork(c *gin.Context) {
	fmt.Println("Welcome to controllers.CreateLiteraryWork")

	var input database.LiteraryWorkInput

	// Bind the JSON request body to the LiteraryWorkInput struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Call the model function to create a new literary work
	newLiteraryWork, err := database.CreateLiteraryWork(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create literary work"})
		return
	}

	c.JSON(http.StatusCreated, newLiteraryWork)
}

// GetLiteraryWorks returns a list of all literary works from the database
func GetLiteraryWorks(c *gin.Context) {
	fmt.Println("Welcome to controllers.GetLiteraryWorks")

	// Call the model function to get all literary works
	literaryWorks, err := database.GetLiteraryWorks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch literary works"})
		return
	}

	c.JSON(http.StatusOK, literaryWorks)
}

// GetLiteraryWorkByID returns a literary work with the specified ID from the database
func GetLiteraryWorkByID(c *gin.Context) {
	fmt.Println("Welcome to controllers.GetLiteraryWorkByID")

	literaryWorkID := c.Param("id")

	// Convert the literaryWorkID to uint
	id, err := strconv.ParseUint(literaryWorkID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid literary work ID"})
		return
	}

	// Call the model function to get the literary work by ID
	literaryWork, err := database.GetLiteraryWorkByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Literary work not found"})
		return
	}

	c.JSON(http.StatusOK, literaryWork)
}

// UpdateLiteraryWork updates an existing literary work in the database
func UpdateLiteraryWork(c *gin.Context) {
	fmt.Println("Welcome to controllers.UpdateLiteraryWork")

	literaryWorkID := c.Param("id")

	// Convert the literaryWorkID to uint
	id, err := strconv.ParseUint(literaryWorkID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid literary work ID"})
		return
	}

	var input database.LiteraryWorkInput

	// Bind the JSON request body to the LiteraryWorkInput struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Call the model function to update the literary work
	updatedLiteraryWork, err := database.UpdateLiteraryWork(uint(id), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update literary work"})
		return
	}

	c.JSON(http.StatusOK, updatedLiteraryWork)
}

// DeleteLiteraryWork deletes a literary work with the specified ID from the database
func DeleteLiteraryWork(c *gin.Context) {
	fmt.Println("Welcome to controllers.DeleteLiteraryWork")

	literaryWorkID := c.Param("id")

	// Convert the literaryWorkID to uint
	id, err := strconv.ParseUint(literaryWorkID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid literary work ID"})
		return
	}

	// Call the model function to delete the literary work
	err = database.DeleteLiteraryWork(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete literary work"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
