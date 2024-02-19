package books

import (
	"database/sql"
	"net/http"
)

func Routes(db *sql.DB) {
	// Define routes
	booksController := BooksController{DB: db}

	http.HandleFunc("/books", booksController.GetBooks)
	http.HandleFunc("/books/", booksController.GetBookByID)
}
