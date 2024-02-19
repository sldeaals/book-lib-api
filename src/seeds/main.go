package seeds

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DefaultBooks contains a list of default book titles for seed data
var DefaultBooks = []string{
	"The Catcher in the Rye",
	"To Kill a Mockingbird",
	"1984",
	"The Great Gatsby",
	"One Hundred Years of Solitude",
	"Brave New World",
	"The Hobbit",
	"The Lord of the Rings",
	"Fahrenheit 451",
	"The Odyssey",
}

// SeedData generates and inserts seed data into the books and pages tables
func SeedData(db *sql.DB) error {
	// Check if the books table is empty
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if count != 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	// Seed data for books
	for _, title := range DefaultBooks {
		// Insert each book into the books table
		result, err := db.Exec("INSERT INTO books (title) VALUES (?)", title)
		if err != nil {
			return err
		}

		// Get the ID of the inserted book
		bookID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		// Generate and insert random pages for each book
		if err := generateAndInsertRandomPages(db, bookID); err != nil {
			return err
		}
	}

	return nil
}

// generateAndInsertRandomPages generates and inserts random pages for a given book ID
func generateAndInsertRandomPages(db *sql.DB, bookID int64) error {
	numPages := rand.Intn(10) + 1 // Random number of pages between 1 and 10
	for i := 1; i <= numPages; i++ {
		// Generate random content for each page (limited to 150 characters)
		content := generateRandomContent(150)

		// Insert each page into the pages table
		_, err := db.Exec("INSERT INTO book_page (book_id, page_num, content) VALUES (?, ?, ?)", bookID, i, content)
		if err != nil {
			return err
		}
	}
	return nil
}

// generateRandomContent generates random content for a page
func generateRandomContent(maxLength int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[{]};:'\",<.>/?"
	content := make([]byte, rand.Intn(maxLength)+1)
	for i := range content {
		content[i] = charset[rand.Intn(len(charset))]
	}
	return string(content)
}
