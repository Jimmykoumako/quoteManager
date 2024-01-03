// File: routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"api/controllers"
	"api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Quotes endpoints
	quoteGroup := r.Group("/api/quotes")
	quoteGroup.Use(middleware.AuthMiddleware())
	{
		quoteGroup.GET("/", controllers.GetQuotes)
		quoteGroup.GET("/:id", controllers.GetQuoteByID)
		quoteGroup.POST("/", controllers.AddQuote)
		quoteGroup.PUT("/:id", controllers.UpdateQuote)
		quoteGroup.DELETE("/:id", controllers.DeleteQuote)
	}

	// User routes
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", middleware.AuthMiddleware(), controllers.LoginUser) // Add middleware here
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		// Add other user-related routes as needed

		// Example: Get user's quotes
		userGroup.GET("/:id/quotes", controllers.GetUserQuotes)

		// Example: Get user's folders
		userGroup.GET("/:id/folders", controllers.GetUserFolders)
	}

	// Feedback endpoints
	feedbackGroup := r.Group("/api/feedback")
	feedbackGroup.Use(middleware.AuthMiddleware())
	{
		feedbackGroup.GET("/:id", controllers.GetFeedbackByID)
		feedbackGroup.GET("/", controllers.GetAllFeedback)
		feedbackGroup.POST("/:quoteId", controllers.AddFeedbackForQuote)
		// Add other feedback-related routes as needed
	}

	// Other endpoints go here...

	return r
}
