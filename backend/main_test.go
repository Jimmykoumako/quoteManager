// main_test.go
package main

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/api/quotes/", controllers.GetQuotes)
	r.GET("/api/quotes/:id", controllers.GetQuoteByID)
	r.POST("/api/quotes/", controllers.AddQuote)
	r.PUT("/api/quotes/:id", controllers.UpdateQuote)
	r.DELETE("/api/quotes/:id", controllers.DeleteQuote)

	r.POST("/api/users/register", controllers.RegisterUser)
	r.POST("/api/users/login", controllers.LoginUser)
	r.GET("/api/users/:id", controllers.GetUserByID)
	r.PUT("/api/users/:id", controllers.UpdateUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)
	r.GET("/api/users/:id/quotes", controllers.GetUserQuotes)
	r.GET("/api/users/:id/folders", controllers.GetUserFolders)

	r.GET("/api/feedback/:id", controllers.GetFeedbackByID)
	r.GET("/api/feedback/", controllers.GetAllFeedback)
	r.POST("/api/feedback/:quoteId", controllers.AddFeedbackForQuote)

	return r
}

func TestRoutes(t *testing.T) {
	router := setupRouter()

	testCases := []struct {
		method string
		path   string
		status int
	}{
		{method: "GET", path: "/api/quotes/", status: http.StatusOK},
		{method: "GET", path: "/api/quotes/123", status: http.StatusOK}, // Update with valid quote ID
		{method: "POST", path: "/api/quotes/", status: http.StatusOK},
		{method: "PUT", path: "/api/quotes/123", status: http.StatusOK},    // Update with valid quote ID
		{method: "DELETE", path: "/api/quotes/123", status: http.StatusOK}, // Update with valid quote ID

		{method: "POST", path: "/api/users/register", status: http.StatusOK},
		{method: "POST", path: "/api/users/login", status: http.StatusOK},
		{method: "GET", path: "/api/users/123", status: http.StatusOK},         // Update with valid user ID
		{method: "PUT", path: "/api/users/123", status: http.StatusOK},         // Update with valid user ID
		{method: "DELETE", path: "/api/users/123", status: http.StatusOK},      // Update with valid user ID
		{method: "GET", path: "/api/users/123/quotes", status: http.StatusOK},  // Update with valid user ID
		{method: "GET", path: "/api/users/123/folders", status: http.StatusOK}, // Update with valid user ID

		{method: "GET", path: "/api/feedback/123", status: http.StatusOK}, // Update with valid feedback ID
		{method: "GET", path: "/api/feedback/", status: http.StatusOK},
		{method: "POST", path: "/api/feedback/456", status: http.StatusOK}, // Update with valid quote ID
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			if rr.Code != tc.status {
				t.Errorf("expected status %d, got %d", tc.status, rr.Code)
			}
		})
	}
}
