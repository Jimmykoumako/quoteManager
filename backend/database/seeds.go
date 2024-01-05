package database

import (
	"errors"
	"api/models"
	"github.com/lib/pq"
)

// SeedDatabase populates the database with additional sample data
func SeedDatabase() error {
	// Categories
	categories := []models.Category{
		{Name: "Philosophy"},
		{Name: "Science"},
		{Name: "Fiction"},
		{Name: "History"},
		{Name: "Technology"},
		{Name: "Self-Help"},
		{Name: "Nature"},
		{Name: "Motivation"},
		{Name: "Art"},
		{Name: "Music"},
		{Name: "Travel"},
		{Name: "Humor"},
		{Name: "Business"},
		{Name: "Health"},
		{Name: "Fitness"},
		{Name: "Cooking"},
		{Name: "Education"},
		{Name: "Sports"},
		{Name: "Love"},
		{Name: "Friendship"},
		{Name: "Movies"},
		{Name: "Spirituality"},
		{Name: "Psychology"},
		{Name: "Economics"},
		{Name: "Politics"},
		{Name: "Space"},
		{Name: "Mythology"},
		{Name: "Gaming"},
		{Name: "Programming"},
		{Name: "Fashion"},
		{Name: "Drama"},
		{Name: "Adventure"},
		{Name: "Thriller"},
		{Name: "Romance"},
		{Name: "Comedy"},
	}

	// Users
	users := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
		{Username: "user3", Password: "password3"},
		{Username: "user4", Password: "password4"},
		{Username: "user5", Password: "password5"},
		// Add more users as needed
	}

	// Folders
	folders := []models.Folder{
		{Name: "Favorites", UserID: 1},
		{Name: "Inspiration", UserID: 2},
		{Name: "Bookmarks", UserID: 3},
		{Name: "Quotes to Remember", UserID: 4},
		{Name: "Daily Motivation", UserID: 5},
		// Add more folders as needed
	}

	// LiteraryWorks
	literaryWorks := []models.LiteraryWork{
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger"},
		{Title: "1984", Author: "George Orwell"},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee"},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
		{Title: "Moby-Dick", Author: "Herman Melville"},
		{Title: "Pride and Prejudice", Author: "Jane Austen"},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien"},
		{Title: "The Art of War", Author: "Sun Tzu"},
		{Title: "The Alchemist", Author: "Paulo Coelho"},
		{Title: "The Odyssey", Author: "Homer"},
		// Add more literary works as needed
	}

	// Quotes
	quotes := []models.Quote{
		{Text: "In the beginning, God created the heavens and the earth.", Author: "Bible", Category: "Religion", UserID: 1, WorkID: 1},
		{Text: "The only thing we have to fear is fear itself.", Author: "Franklin D. Roosevelt", Category: "History", UserID: 2, WorkID: 2},
		{Text: "To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.", Author: "Ralph Waldo Emerson", Category: "Motivation", UserID: 3, WorkID: 3},
		{Text: "The greatest glory in living lies not in never falling, but in rising every time we fall.", Author: "Nelson Mandela", Category: "Inspiration", UserID: 4, WorkID: 4},
		{Text: "Life is what happens when you're busy making other plans.", Author: "John Lennon", Category: "Life", UserID: 5, WorkID: 5},
		{Text: "To be or not to be, that is the question.", Author: "William Shakespeare", Category: "Drama", UserID: 5, WorkID: 6},
		{Text: "The only limit to our realization of tomorrow will be our doubts of today.", Author: "Franklin D. Roosevelt", Category: "Inspiration", UserID: 1, WorkID: 7},
		{Text: "It does not matter how slowly you go as long as you do not stop.", Author: "Confucius", Category: "Motivation", UserID: 2, WorkID: 8},
		{Text: "The only way to do great work is to love what you do.", Author: "Steve Jobs", Category: "Technology", UserID: 3, WorkID: 9},
		{Text: "Success is not final, failure is not fatal: It is the courage to continue that counts.", Author: "Winston Churchill", Category: "Motivation", UserID: 4, WorkID: 10},
		
		// Add more quotes as needed
	}

	// Likes
	likes := []models.Like{
		{UserID: 1, QuoteID: 31, IsLike: true},
		{UserID: 2, QuoteID: 31, IsLike: true},
		{UserID: 3, QuoteID: 32, IsLike: true},
		{UserID: 4, QuoteID: 33, IsLike: true},
		{UserID: 5, QuoteID: 34, IsLike: true},
		// Add more likes as needed
	}
	
	// Feedbacks
	feedbacks := []models.Feedback{
		{Comment: "Great quote!", Rating: 5, QuoteID: 1},
		{Comment: "Interesting perspective.", Rating: 4, QuoteID: 2},
		{Comment: "Well said!", Rating: 5, QuoteID: 3},
		{Comment: "I disagree.", Rating: 2, QuoteID: 4},
		// Add more feedbacks as needed
	}
	// Create records in the database
	
	if err := createWithSkipDuplicates(&categories); err != nil {
		return err
	}

	if err := createWithSkipDuplicates(&users); err != nil {
		return err
	}

	if err := createWithSkipDuplicates(&folders); err != nil {
		return err
	}
	
	if err := createWithSkipDuplicates(&literaryWorks); err != nil {
		return err
	}
	
	if err := createWithSkipDuplicates(&quotes); err != nil {
		return err
	}

	if err := createWithSkipDuplicates(&likes); err != nil {
		return err
	}
	
	if err := createWithSkipDuplicates(&feedbacks); err != nil {
		return err
	}

	return nil
}

// createWithSkipDuplicates creates records in the database and skips duplicate errors
func createWithSkipDuplicates(data interface{}) error {
	if err := db.Create(data).Error; err != nil {
		// Check if the error is a duplicate key violation
		if isDuplicateKeyError(err) {
			return nil // Skip insertion on duplicate key violation
		}
		return err
	}
	return nil
}

// isDuplicateKeyError checks if the error is a duplicate key violation
func isDuplicateKeyError(err error) bool {
	var pqError *pq.Error
	if errors.As(err, &pqError) && pqError.Code == "23505" {
		return true
	}
	return false
}

