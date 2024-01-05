// routes/routes_test.go
package routes

import (
	"api/database"
	"api/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Define test input
	payload := `{"username": "testuser", "password": "testpassword"}`

	// Create a request
	req, err := http.NewRequest("POST", "/api/public/users/register", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `{"message": "User registered successfully"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	database.CloseDB()
}

func TestGetQuotes(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a request
	req, err := http.NewRequest("GET", "/api/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestAddQuote(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Define test input
	payload := `{"text": "Test quote", "author": "Test author", "category": "Test category", "user_id": 1}`

	// Create a request
	req, err := http.NewRequest("POST", "/api/quotes", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestAddFeedbackForQuote(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a test quote
	quote := models.Quote{
		Text:     "Test quote for feedback",
		Author:   "Test author",
		Category: "Test category",
		UserID:   1,
	}
	// Get the database reference
	db := database.GetDB()
	db.Create(&quote)

	// Define test input
	payload := `{"comment": "Great quote!", "rating": 5}`

	// Create a request
	req, err := http.NewRequest("POST", "/api/feedback/1", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

// routes/routes_test.go (continued)

func TestGetUserByID(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a request
	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestUpdateUser(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Define test input
	payload := `{"username": "updateduser", "password": "updatedpassword"}`

	// Create a request
	req, err := http.NewRequest("PUT", "/api/users/1", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestDeleteUser(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a request
	req, err := http.NewRequest("DELETE", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestGetUserQuotes(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a request
	req, err := http.NewRequest("GET", "/api/users/1/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestGetUserFolders(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a request
	req, err := http.NewRequest("GET", "/api/users/1/folders", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestGetFeedbackByID(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a test feedback
	feedback := models.Feedback{
		Comment: "Test feedback",
		Rating:  5,
		QuoteID: 1,
	}
	// Get the database reference
	db := database.GetDB()
	db.Create(&feedback)

	// Create a request
	req, err := http.NewRequest("GET", "/api/feedback/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}

func TestGetAllFeedback(t *testing.T) {
	r := SetupRouter()
	database.InitDB()

	// Create a request
	req, err := http.NewRequest("GET", "/api/feedback", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// You can add more assertions based on your controller's logic and expected behavior

	database.CloseDB()
}
