module github.com/book-lib-api

go 1.21.7

replace github.com/book-lib-api/src/database => ../database

replace github.com/book-lib-api/src/seeds => ../seeds

replace github.com/book-lib-api/src/structs => ../structs

replace github.com/book-lib-api/src/books => ../books

replace github.com/book-lib-api/src/book_page => ../book_page

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
