package book_page

// Page struct represents a page in a book
type BookPage struct {
	BookID  int    `json:"book_id"`
	PageNum int    `json:"page_num"`
	Content string `json:"content"`
}
