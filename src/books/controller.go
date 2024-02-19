package books

import (
	"database/sql"
	"net/http"
)

// BooksController handles requests for books routes
type BooksController struct {
	DB *sql.DB
}

// GetBooks returns the list of books
func (c *BooksController) GetBooks(w http.ResponseWriter, r *http.Request) {
	Books(w, r, c.DB)
}

// GetBooksByID returns a single book
func (c *BooksController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	BookByID(w, r, c.DB)
}
