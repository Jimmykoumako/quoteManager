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

	// Apply global middlewares
	applyGlobalMiddlewares(r)

	// Define public routes (no authentication required)
	definePublicRoutes(r)

	// Define protected routes (authentication required)
	defineProtectedRoutes(r)

	return r
}

// applyGlobalMiddlewares applies global middlewares to the router
func applyGlobalMiddlewares(r *gin.Engine) {
	// Apply tracing middleware to all routes
	r.Use(middleware.TracingMiddleware())

	// Apply security headers middleware to all routes
	r.Use(middleware.SecurityHeadersMiddleware())

	// Apply request logging middleware to all routes
	r.Use(middleware.LogRequest())

	// Create a rate limiter (example: 20 requests per second)
	// limiter := rate.NewLimiter(rate.Limit(20), 1)

	// Apply rate limiting middleware to all routes
	r.Use(middleware.RateLimiterMiddleware(limiter))

	// Apply log feedback action middleware to specific routes
	r.Use(middleware.LogFeedbackAction())

	// Apply error handling middleware to all routes
	r.Use(middleware.ErrorHandlerMiddleware())
}

// definePublicRoutes defines routes that do not require authentication
func definePublicRoutes(r *gin.Engine) {
	publicGroup := r.Group("/api/public")
	{
		publicGroup.POST("/users/register", controllers.RegisterUser)
		publicGroup.POST("/users/login", controllers.LoginUser)
		// Add other public routes as needed
	}
}

// defineProtectedRoutes defines routes that require authentication
func defineProtectedRoutes(r *gin.Engine) {
	protectedGroup := r.Group("/api")
	protectedGroup.Use(middleware.AuthMiddleware())            // Apply authentication middleware to all routes in this group
	protectedGroup.Use(middleware.TracingMiddleware())         // Apply tracing middleware to all routes in this group
	protectedGroup.Use(middleware.SecurityHeadersMiddleware()) // Apply security headers middleware to all routes in this group
	protectedGroup.Use(middleware.LogRequest())                // Apply request logging middleware to all routes in this group
	//protectedGroup.Use(middleware.RateLimiterMiddleware(limiter)) // Apply rate limiting middleware to all routes in this group
	protectedGroup.Use(middleware.LogFeedbackAction())      // Apply log feedback action middleware to all routes in this group
	protectedGroup.Use(middleware.ErrorHandlerMiddleware()) // Apply error handling middleware to all routes in this group

	// Define routes for different entities
	defineQuoteRoutes(protectedGroup)
	defineUserRoutes(protectedGroup)
	defineFeedbackRoutes(protectedGroup)
	defineFolderRoutes(protectedGroup)
	defineLikeRoutes(r) // Note: Not added to the protectedGroup as it seems to be separate
	defineCategoryRoutes(protectedGroup)
	defineTagRoutes(protectedGroup)
	// Add other protected endpoints go here...
}

// defineQuoteRoutes defines routes related to quotes
func defineQuoteRoutes(protectedGroup *gin.RouterGroup) {
	quoteGroup := protectedGroup.Group("/quotes")
	{
		quoteGroup.GET("/", controllers.GetQuotes)
		quoteGroup.GET("/:id", controllers.GetQuoteByID)
		quoteGroup.POST("/", controllers.AddQuote)
		quoteGroup.PUT("/:id", controllers.UpdateQuote)
		quoteGroup.DELETE("/:id", controllers.DeleteQuote)
	}
}

// defineUserRoutes defines routes related to users
func defineUserRoutes(protectedGroup *gin.RouterGroup) {
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
}

// defineFeedbackRoutes defines routes related to feedback
func defineFeedbackRoutes(protectedGroup *gin.RouterGroup) {
	feedbackGroup := protectedGroup.Group("/feedback")
	{
		feedbackGroup.GET("/:id", controllers.GetFeedbackByID)
		feedbackGroup.GET("/", controllers.GetAllFeedback)
		feedbackGroup.POST("/:quoteId", controllers.AddFeedbackForQuote)
		feedbackGroup.PUT("/:id", controllers.UpdateFeedback)
		feedbackGroup.DELETE("/:id", controllers.DeleteFeedback)
		// Add other feedback-related routes as needed
	}
}

// defineFolderRoutes defines routes related to folders
func defineFolderRoutes(protectedGroup *gin.RouterGroup) {
	folderGroup := protectedGroup.Group("/folders")
	{
		folderGroup.GET("/:id", controllers.GetFolderByID)
		folderGroup.GET("/", controllers.GetFoldersForUser)
		folderGroup.POST("/", controllers.CreateFolder)
		folderGroup.PUT("/:id", controllers.UpdateFolder)
		folderGroup.DELETE("/:id", controllers.DeleteFolder)
		// Add other folder-related routes as needed
	}
}

// defineLikeRoutes defines routes related to likes
func defineLikeRoutes(r *gin.Engine) {
	likeGroup := r.Group("/api/likes")
	{
		likeGroup.GET("/:id", controllers.GetLikeByID)
		likeGroup.POST("/", controllers.AddLike)
		likeGroup.PUT("/:id", controllers.UpdateLike)
		likeGroup.DELETE("/:id", controllers.DeleteLike)
	}
}

// defineCategoryRoutes defines routes related to categories
func defineCategoryRoutes(protectedGroup *gin.RouterGroup) {
    categoryGroup := protectedGroup.Group("/categories")
    {
        categoryGroup.GET("/", controllers.GetCategories)
        categoryGroup.GET("/:id", controllers.GetCategoryByID)
        categoryGroup.POST("/", controllers.CreateCategory)
        categoryGroup.PUT("/:id", controllers.UpdateCategory)
        categoryGroup.DELETE("/:id", controllers.DeleteCategory)
    }
}

// defineTagRoutes defines routes related to tags
func defineTagRoutes(protectedGroup *gin.RouterGroup) {
	tagGroup := protectedGroup.Group("/tags")
	{
		tagGroup.GET("/", controllers.GetTags)
		tagGroup.GET("/:id", controllers.GetTagByID)
		tagGroup.POST("/", controllers.CreateTag)
		tagGroup.PUT("/:id", controllers.UpdateTag)
		tagGroup.DELETE("/:id", controllers.DeleteTag)
	}
}
