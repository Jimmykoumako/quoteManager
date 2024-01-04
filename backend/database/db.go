package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// SQL statements for table creation
const (
	createUsersTable = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`

	createQuotesTable = `
		CREATE TABLE IF NOT EXISTS quotes (
			id SERIAL PRIMARY KEY,
			text TEXT NOT NULL,
			author VARCHAR(255),
			source_id INT,
			category_id INT,
			creation_date TIMESTAMP,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`

	createCategoriesTable = `
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			description TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`

	createSourcesTable = `
		CREATE TABLE IF NOT EXISTS sources (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			description TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`

	createCommentsTable = `
		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			user_id INT,
			quote_id INT,
			text TEXT NOT NULL,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`

	createFavoritesTable = `
		CREATE TABLE IF NOT EXISTS favorites (
			id SERIAL PRIMARY KEY,
			user_id INT,
			quote_id INT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		)
	`

	createNotificationsTable = `
		CREATE TABLE IF NOT EXISTS notifications (
			id SERIAL PRIMARY KEY,
			user_id INT,
			message TEXT NOT NULL,
			created_at TIMESTAMP
		)
	`

	createActivityLogsTable = `
		CREATE TABLE IF NOT EXISTS activity_logs (
			id SERIAL PRIMARY KEY,
			user_id INT,
			action VARCHAR(255) NOT NULL,
			created_at TIMESTAMP
		)
	`

	createFeedbackTable = `
		CREATE TABLE IF NOT EXISTS feedback (
			id SERIAL PRIMARY KEY,
			user_id INT,
			message TEXT NOT NULL,
			created_at TIMESTAMP
		)
	`
)

// Initialize initializes the database connection.
func Initialize(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables if they don't exist
	err = createTables(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// createTables creates the required tables if they don't exist
func createTables(db *sql.DB) error {
	// Use a transaction to execute multiple SQL statements
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-panic after rollback
		} else if err != nil {
			_ = tx.Rollback() // rollback if there was an error
		} else {
			err = tx.Commit() // commit if there was no error
		}
	}()

	// Execute the SQL statements
	_, err = tx.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	_, err = tx.Exec(createQuotesTable)
	if err != nil {
		return fmt.Errorf("failed to create quotes table: %w", err)
	}

	_, err = tx.Exec(createCategoriesTable)
	if err != nil {
		return fmt.Errorf("failed to create categories table: %w", err)
	}

	_, err = tx.Exec(createSourcesTable)
	if err != nil {
		return fmt.Errorf("failed to create sources table: %w", err)
	}

	_, err = tx.Exec(createCommentsTable)
	if err != nil {
		return fmt.Errorf("failed to create comments table: %w", err)
	}

	_, err = tx.Exec(createFavoritesTable)
	if err != nil {
		return fmt.Errorf("failed to create favorites table: %w", err)
	}

	_, err = tx.Exec(createNotificationsTable)
	if err != nil {
		return fmt.Errorf("failed to create notifications table: %w", err)
	}

	_, err = tx.Exec(createActivityLogsTable)
	if err != nil {
		return fmt.Errorf("failed to create activity logs table: %w", err)
	}

	_, err = tx.Exec(createFeedbackTable)
	if err != nil {
		return fmt.Errorf("failed to create feedback table: %w", err)
	}

	return nil
}
