package book_page

import (
	"database/sql"
	"net/http"
)

func Routes(db *sql.DB) {
	// Define routes
	bookPageController := BookPageController{DB: db}

	http.HandleFunc("/book/", bookPageController.GetBookPageByID)
}
