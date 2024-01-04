// controllers/feedback_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"api/models"
	"strconv"
	"math/rand"
	"fmt"
)

// In-memory storage for demonstration purposes
var feedbackStorage = make(map[string]models.Feedback)

// GetFeedbackByID returns a specific feedback by ID
func GetFeedbackByID(c *gin.Context) {
	// Extract feedback ID from the URL parameter
	feedbackID := c.Param("id")

	// Fetch feedback from the in-memory storage
	feedback, found := feedbackStorage[feedbackID]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

// GetAllFeedback returns all feedback entries
func GetAllFeedback(c *gin.Context) {
	// Convert feedback map to a slice
	var feedbackList []models.Feedback
	for _, v := range feedbackStorage {
		feedbackList = append(feedbackList, v)
	}

	c.JSON(http.StatusOK, feedbackList)
}

// AddFeedbackForQuote adds feedback for a specific quote
func AddFeedbackForQuote(c *gin.Context) {
	// Extract quote ID from the URL parameter
	quoteID := c.Param("quoteId")

	// Implement logic to add feedback for a quote to the in-memory storage
	var feedback models.Feedback

	// Bind the request body to the feedback model
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate feedback data
	if feedback.Rating < 1 || feedback.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating. Rating must be between 1 and 5."})
		return
	}

	if len(feedback.Comment) > 500 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment is too long. Maximum length is 500 characters."})
		return
	}

	if feedback.Comment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment cannot be empty."})
		return
	}

	// Set the QuoteID based on the parameter
	// Convert string to uint
	quoteIDUint, err := strconv.ParseUint(quoteID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID."})
		return
	}
	feedback.QuoteID = uint(quoteIDUint)

	// Save feedback to the in-memory storage
	feedbackID := generateFeedbackID()
	feedbackStorage[feedbackID] = feedback

	// Parse feedbackID to uint and assign it to feedback.ID
	parsedID, err := strconv.ParseUint(feedbackID, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	feedback.ID = uint(parsedID)

	c.JSON(http.StatusCreated, feedback)
}
// generateFeedbackID generates a unique ID for feedback (for demonstration purposes)
func generateFeedbackID() string {
	// In a real-world scenario, you would use a more sophisticated method to generate unique IDs
	return "fb" + randomString(6)
}

// randomString generates a random string of a specified length (for demonstration purposes)
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
