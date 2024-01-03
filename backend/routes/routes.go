package routes

import (
	"github.com/gin-gonic/gin"
	"api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Quotes endpoints
	quoteGroup := r.Group("/api/quotes")
	quoteGroup.Use(middleware.AuthMiddleware())
    {
        quoteRoutes.GET("/", controllers.GetQuotes)
        quoteRoutes.GET("/:id", controllers.GetQuoteByID)
        quoteRoutes.POST("/", controllers.AddQuote)
        quoteRoutes.PUT("/:id", controllers.UpdateQuote)
        quoteRoutes.DELETE("/:id", controllers.DeleteQuote)
    }
	
	// User routes
	userGroup := r.Group("/api/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
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
