package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"api/models"
	"api/database"
	"strconv"
)

// GetQuotes returns all quotes
func GetQuotes(c *gin.Context) {
	quotes, err := database.GetQuotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quotes"})
		return
	}

	c.JSON(http.StatusOK, quotes)
}

// GetQuoteByID returns a specific quote by ID
func GetQuoteByID(c *gin.Context) {
	quoteID := c.Param("id")

	// Convert quoteID to uint
	quoteIDUint, err := convertToUint(quoteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	quote, err := database.GetQuoteByID(quoteIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the quote"})
		return
	}

	c.JSON(http.StatusOK, quote)
}

// AddQuote adds a new quote
func AddQuote(c *gin.Context) {
	var newQuote models.Quote

	// Bind JSON request body to Quote model
	if err := c.ShouldBindJSON(&newQuote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Add the new quote to the database
	addedQuote, err := database.AddQuote(newQuote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add the quote"})
		return
	}

	c.JSON(http.StatusCreated, addedQuote)
}

// UpdateQuote updates an existing quote by ID
func UpdateQuote(c *gin.Context) {
	quoteID := c.Param("id")

	// Convert quoteID to uint
	quoteIDUint, err := convertToUint(quoteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	var updatedQuote models.Quote

	// Bind JSON request body to updated Quote model
	if err := c.ShouldBindJSON(&updatedQuote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update the quote in the database
	updatedQuote, err = database.UpdateQuote(quoteIDUint, updatedQuote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the quote"})
		return
	}

	c.JSON(http.StatusOK, updatedQuote)
}

// DeleteQuote deletes a quote by ID
func DeleteQuote(c *gin.Context) {
	quoteID := c.Param("id")

	// Convert quoteID to uint
	quoteIDUint, err := convertToUint(quoteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	// Delete the quote from the database
	err = database.DeleteQuote(quoteIDUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the quote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quote deleted successfully"})
}


// convertToUint converts a string to uint.
// It returns an error if the conversion fails or if the result is negative.
func convertToUint(s string) (uint, error) {
	converted, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if converted < 0 {
		return 0, fmt.Errorf("negative value not allowed: %d", converted)
	}

	return uint(converted), nil
}

