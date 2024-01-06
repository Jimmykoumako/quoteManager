// database/user_profile_db.go

package database

import (
    "api/models"
)

// GetUserProfileByID retrieves a user profile by ID
func GetUserProfileByID(profileID uint) (*models.UserProfile, error) {
    var userProfile models.UserProfile
    if err := db.First(&userProfile, profileID).Error; err != nil {
        return nil, err
    }
    return &userProfile, nil
}

// CreateUserProfile creates a new user profile
func CreateUserProfile(profile *models.UserProfile) error {
    if err := db.Create(profile).Error; err != nil {
        return err
    }
    return nil
}

// UpdateUserProfile updates an existing user profile by ID
func UpdateUserProfile(profileID uint, updatedProfile *models.UserProfile) error {
    if err := db.Model(&models.UserProfile{}).Where("id = ?", profileID).Updates(updatedProfile).Error; err != nil {
        return err
    }
    return nil
}

// DeleteUserProfile deletes a user profile by ID
func DeleteUserProfile(profileID uint) error {
    if err := db.Where("id = ?", profileID).Delete(&models.UserProfile{}).Error; err != nil {
        return err
    }
    return nil
}
