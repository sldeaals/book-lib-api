package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Config holds the database configuration
type Config struct {
	User     string
	Password string
	Network  string
	Address  string
	Port     string
	Name     string
}

// InitDB initializes the database connection and schema
func InitDB(config Config) (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", config.User, config.Password, config.Network, config.Address, config.Port, config.Name)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Ensure the database is reachable
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create the database if it does not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Name)
	if err != nil {
		return nil, err
	}

	// Switch to the database
	_, err = db.Exec("USE " + config.Name)
	if err != nil {
		return nil, err
	}

	// Create the books table if it does not exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255) NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	// Create the pages table if it does not exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS book_page (
			book_id INT,
			page_num INT,
			content TEXT,
			PRIMARY KEY (book_id, page_num),
			FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
