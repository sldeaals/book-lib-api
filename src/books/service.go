package books

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/book-lib-api/src/structs"
)

// BooksResponse is the response structure for the Books API
type BooksResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Payload []Book `json:"data,omitempty"`
}

// BookResponse is the response structure for a single book
type BookResponse struct {
	Code    int  `json:"code"`
	Status  string `json:"status"`
	Payload Book   `json:"data,omitempty"`
}

// Books handles the API endpoint to get a list of books
func Books(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	books, err := getBooks(db)
	if err != nil {
		structs.HandleError(w, r, structs.NewError(http.StatusInternalServerError, "Internal Error", err))
		return
	}

	response := BooksResponse{Code: http.StatusOK, Status: "Success", Payload: books}
	json.NewEncoder(w).Encode(response)
}

// getBooks queries for the current list of books
func getBooks(db *sql.DB) ([]Book, error) {
	var books []Book

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, fmt.Errorf("booksList: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title); err != nil {
			return nil, fmt.Errorf("booksList: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("booksList: %v", err)
	}

	return books, nil
}

// BookByID handles the API endpoint to get a specific book by ID
func BookByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))
	if err != nil {
		structs.HandleError(w, r, structs.NewError(http.StatusBadRequest, "Invalid book ID", err))
		return
	}

	book, err := getBookByID(db, id)
	if err != nil {
		structs.HandleError(w, r, structs.NewError(http.StatusNotFound, "Book not found", err))
		return
	}

	response := BookResponse{Code: http.StatusOK, Status: "Success", Payload: book}
	json.NewEncoder(w).Encode(response)
}

// getBookByID queries for a specific book by ID
func getBookByID(db *sql.DB, id int) (Book, error) {
	var book Book

	row := db.QueryRow("SELECT * FROM books WHERE id = ?", id)
	if err := row.Scan(&book.ID, &book.Title); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("booksById %d: no such book", id)
		}
		return book, fmt.Errorf("booksById %d: %v", id, err)
	}

	return book, nil
}
