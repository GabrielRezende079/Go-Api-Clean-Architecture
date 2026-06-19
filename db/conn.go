package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "go_tests"
)

// ConnerctDB establishes a connection to the PostgreSQL database and returns the connection object.
func ConnerctDB() (*sql.DB, error) {
	// Create the connection string using the provided database credentials
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo) // Open a connection to the database

	// Check for errors while opening the connection
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// If we reach this point, the connection is successful
	fmt.Println("Successfully connected to database!")
	return db, nil
}
