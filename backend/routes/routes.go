// routes/routes.go

package routes

import (
	"api/controllers"
	"api/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Apply tracing middleware to all routes
	r.Use(middleware.TracingMiddleware())

	// Apply security headers middleware to all routes
	r.Use(middleware.SecurityHeadersMiddleware())

	// Apply request logging middleware to all routes
	r.Use(middleware.LogRequest())

	// Create a rate limiter (example: 10 requests per second)
	limiter := rate.NewLimiter(rate.Limit(20), 1)

	// Apply rate limiting middleware to all routes
	r.Use(middleware.RateLimiterMiddleware(limiter))

	// Apply log feedback action middleware to specific routes
	r.Use(middleware.LogFeedbackAction())

	// Apply error handling middleware to all routes
	r.Use(middleware.ErrorHandlerMiddleware())

	// Public routes (no authentication required)
	publicGroup := r.Group("/api/public")
	{
		publicGroup.POST("/users/register", controllers.RegisterUser)
		publicGroup.POST("/users/login", controllers.LoginUser)
		// Add other public routes as needed
	}

	// Protected routes (authentication required)
	protectedGroup := r.Group("/api")
	protectedGroup.Use(middleware.AuthMiddleware())            // Apply authentication middleware to all routes in this group
	protectedGroup.Use(middleware.TracingMiddleware())         // Apply tracing middleware to all routes in this group
	protectedGroup.Use(middleware.SecurityHeadersMiddleware()) // Apply security headers middleware to all routes in this group
	protectedGroup.Use(middleware.LogRequest())                // Apply request logging middleware to all routes in this group
	//protectedGroup.Use(middleware.RateLimiterMiddleware(limiter)) // Apply rate limiting middleware to all routes in this group
	protectedGroup.Use(middleware.LogFeedbackAction())      // Apply log feedback action middleware to all routes in this group
	protectedGroup.Use(middleware.ErrorHandlerMiddleware()) // Apply error handling middleware to all routes in this group
	{
		// Quotes endpoints
		quoteGroup := protectedGroup.Group("/quotes")
		{
			quoteGroup.GET("/", controllers.GetQuotes)
			quoteGroup.GET("/:id", controllers.GetQuoteByID)
			quoteGroup.POST("/", controllers.AddQuote)
			quoteGroup.PUT("/:id", controllers.UpdateQuote)
			quoteGroup.DELETE("/:id", controllers.DeleteQuote)
		}

		// User routes
		userGroup := protectedGroup.Group("/users")
		{
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
		feedbackGroup := protectedGroup.Group("/feedback")
		{
			feedbackGroup.GET("/:id", controllers.GetFeedbackByID)
			feedbackGroup.GET("/", controllers.GetAllFeedback)
			feedbackGroup.POST("/:quoteId", controllers.AddFeedbackForQuote)
			// Add other feedback-related routes as needed
		}

		// Other protected endpoints go here...
	}

	return r
}
