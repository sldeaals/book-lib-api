package book_page

import (
	"database/sql"
	"net/http"
)

// BookPageController handles requests for the /book/ route
type BookPageController struct {
	DB *sql.DB
}

// GetBooks returns the list of books
func (c *BookPageController) GetBookPageByID(w http.ResponseWriter, r *http.Request) {
	BookPageByID(w, r, c.DB)
}

