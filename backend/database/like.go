package database

import (
	"api/models"
)

// GetLikeByID fetches a like by ID from the database
func GetLikeByID(likeID string) (*models.Like, error) {
	var like models.Like
	if err := db.Where("id = ?", likeID).First(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}

// AddLike adds a new like to the database
func AddLike(like *models.Like) error {
	if err := db.Create(like).Error; err != nil {
		return err
	}
	return nil
}

// UpdateLike updates an existing like by ID in the database
func UpdateLike(likeID string, updatedLike *models.Like) error {
	var existingLike models.Like
	if err := db.Where("id = ?", likeID).First(&existingLike).Error; err != nil {
		return err
	}

	// Update fields of the existing like with the values from updatedLike
	existingLike.UserID = updatedLike.UserID
	existingLike.QuoteID = updatedLike.QuoteID
	existingLike.IsLike = updatedLike.IsLike

	if err := db.Save(&existingLike).Error; err != nil {
		return err
	}
	return nil
}

// DeleteLike deletes a like by ID from the database
func DeleteLike(likeID string) error {
	if err := db.Where("id = ?", likeID).Delete(&models.Like{}).Error; err != nil {
		return err
	}
	return nil
}
