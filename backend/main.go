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
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Defer the database closing function
	defer func() {
		if err := database.CloseDB(db); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Initialize your router and set up routes
	router := routes.SetupRouter()

	// Set up server configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// Start the server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server is running on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

  