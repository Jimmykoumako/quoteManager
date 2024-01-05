// controllers/feedback_controller.go

package controllers

import (
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

// In-memory storage for demonstration purposes
var feedbackStorage = make(map[string]models.Feedback)

// GetFeedbackByID returns a specific feedback by ID
func GetFeedbackByID(c *gin.Context) {
	feedbackID := c.Param("id")

	feedback, found := feedbackStorage[feedbackID]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

// GetAllFeedback returns all feedback entries
func GetAllFeedback(c *gin.Context) {
	var feedbackList []models.Feedback
	for _, v := range feedbackStorage {
		feedbackList = append(feedbackList, v)
	}

	c.JSON(http.StatusOK, feedbackList)
}

// AddFeedbackForQuote adds feedback for a specific quote
func AddFeedbackForQuote(c *gin.Context) {
	quoteID := c.Param("quoteId")
	var feedback models.Feedback

	if err := c.ShouldBindJSON(&feedback); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateFeedback(feedback); err != nil {
		log.Printf("Invalid feedback: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feedback.QuoteID = parseQuoteID(quoteID)
	feedbackID := generateFeedbackID()
	parsedID, err := strconv.ParseUint(feedbackID, 10, 64)
	if err != nil {
		log.Printf("Error parsing feedback ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	feedback.ID = uint(parsedID)
	feedbackStorage[feedbackID] = feedback

	log.Printf("Feedback added successfully. ID: %d, QuoteID: %d", feedback.ID, feedback.QuoteID)

	c.JSON(http.StatusCreated, feedback)
}

// UpdateFeedback updates a specific feedback
func UpdateFeedback(c *gin.Context) {
	feedbackID := c.Param("id")
	var updatedFeedback models.Feedback

	if err := c.ShouldBindJSON(&updatedFeedback); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, found := feedbackStorage[feedbackID]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	feedbackStorage[feedbackID] = updatedFeedback

	log.Printf("Feedback updated successfully. ID: %s", feedbackID)

	c.JSON(http.StatusOK, updatedFeedback)
}

// DeleteFeedback deletes a specific feedback by ID
func DeleteFeedback(c *gin.Context) {
	feedbackID := c.Param("id")

	_, found := feedbackStorage[feedbackID]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	delete(feedbackStorage, feedbackID)

	log.Printf("Feedback deleted successfully. ID: %s", feedbackID)

	c.JSON(http.StatusNoContent, nil)
}

// generateFeedbackID generates a unique ID for feedback (for demonstration purposes)
func generateFeedbackID() string {
	// Generate a new UUID
	uuid := uuid.New()

	// Convert UUID to string and return
	return uuid.String()
}

func validateFeedback(feedback models.Feedback) error {
	// Validate rating
	if feedback.Rating < 1 || feedback.Rating > 5 {
		return fmt.Errorf("invalid rating. Rating must be between 1 and 5")
	}

	// Validate comment length
	const maxCommentLength = 200
	if len(feedback.Comment) > maxCommentLength {
		return fmt.Errorf("comment is too long. Maximum length is %d characters", maxCommentLength)
	}

	// Validate non-empty comment
	if feedback.Comment == "" {
		return fmt.Errorf("comment cannot be empty")
	}

	// Add more validation rules as needed...

	return nil
}

func parseQuoteID(quoteID string) uint {
	parsedID, err := strconv.ParseUint(quoteID, 10, 64)
	if err != nil {
		fmt.Println("Error parsing quote ID:", err)
	}
	return uint(parsedID)
}
