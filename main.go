package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/book-lib-api/src/database"
	"github.com/book-lib-api/src/books"
	bookPage "github.com/book-lib-api/src/book_page"
	"github.com/book-lib-api/src/seeds"
	"github.com/book-lib-api/src/structs"
)

var db *sql.DB

func main() {
    // Load environment variables
	structs.LoadEnv()

	// Initialize the database
	if err := initializeDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Upload Initial Data
	if err := seeds.SeedData(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seed data loaded successfully.")

	// Setup routes
	setupRoutes()

	// Start the server
	port, err := getPortFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is running on :%d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func initializeDB() error {
	config := database.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Network:  os.Getenv("DB_NETWORK"),
		Address:  os.Getenv("DB_ADDRESS"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	var err error
	db, err = database.InitDB(config)
	return err
}

func setupRoutes() {
	// Setup routes for books
	books.Routes(db)

	// Setup routes for book pages
	bookPage.Routes(db)
}

func getPortFromEnv() (int, error) {
	portStr := os.Getenv("SERVER_PORT")
	if portStr == "" {
		// Default port if not specified in .env
		return 8080, nil
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert SERVER_PORT to integer: %v", err)
	}

	return port, nil
}
