package book_page

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// BookPageResponse is the response structure for the BookPage API
type BookPageResponse struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Payload BookPage `json:"data,omitempty"`
}

// GetPageHandler returns a specific page of a book in the desired format
func BookPageByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	// Extract book ID, page number, and format from the URL
	parts := strings.Split(r.URL.Path, "/")

	// Handle route parameters correctness
	if len(parts) != 6 {
		handleError(w, r, http.StatusBadRequest, "Invalid route", nil)
		return
	}

	bookID, err := strconv.Atoi(parts[2])
	if err != nil {
		handleError(w, r, http.StatusBadRequest, "Invalid book ID", err)
		return
	}

	pageNum, err := strconv.Atoi(parts[4])
	if err != nil {
		handleError(w, r, http.StatusBadRequest, "Invalid page number", err)
		return
	}

	bookPage, err := getBookPageByID(db, bookID, pageNum)
	if err != nil {
		handleError(w, r, http.StatusNotFound, "Book Page not found", err)
		return
	}

	// Handle different formats
	format := parts[5]
	handleFormat(w, bookPage, format)
}

// getBookPageByID retrieves a book page by book ID and page number
func getBookPageByID(db *sql.DB, bookID int, pageNum int) (BookPage, error) {
	var bookPage BookPage

	row := db.QueryRow("SELECT book_id, page_num, content FROM book_page WHERE book_id = ? AND page_num = ?", bookID, pageNum)
	if err := row.Scan(&bookPage.BookID, &bookPage.PageNum, &bookPage.Content); err != nil {
		if err == sql.ErrNoRows {
			return bookPage, fmt.Errorf("bookPageById %d: no such book page", pageNum)
		}
		return bookPage, fmt.Errorf("bookPageById %d: %v", pageNum, err)
	}
	return bookPage, nil
}

// handleError sends an HTTP error response with a standard error message
func handleError(w http.ResponseWriter, r *http.Request, statusCode int, message string, err error) {
	http.Error(w, message, statusCode)
}

// handleFormat sets the appropriate content type and writes the content to the response writer
func handleFormat(w http.ResponseWriter, bookPage BookPage, format string) {
	if format == "html" {
		w.Header().Set("Content-Type", "text/html")
		// Convert plain text to HTML for simplicity
		fmt.Fprintf(w, "<html><body>%s</body></html>", bookPage.Content)
	} else {
		fmt.Fprint(w, bookPage.Content)
	}
}
