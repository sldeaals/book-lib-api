module github.com/book-lib-api

go 1.21.7

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/sys v0.17.0 // indirect
)

replace github.com/book-lib-api/src/database => ../database

replace github.com/book-lib-api/src/seeds => ../seeds

replace github.com/book-lib-api/src/structs => ../structs

replace github.com/book-lib-api/src/books => ../books

replace github.com/book-lib-api/src/book_page => ../book_page
