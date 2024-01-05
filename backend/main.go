// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"api/database"
	"api/routes"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Defer the database closing function
	defer func() {
		database.CloseDB()
	}()

	// // Seed data
	// if err := database.SeedDatabase(); err != nil {
	// 	log.Fatalf("Error seeding data: %v", err)
	// }

	// Initialize your router and set up routes with middleware
	router := routes.SetupRouter()

	// Set up server configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // default port
	}

	// Start the server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
