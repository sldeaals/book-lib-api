# Book Library API (book-lib-api)
Golang book library API

Online library. REST API which allows clients to consume the list of available books, as well as to read those books page by page in the desired formats.

Books will be available (page by page) in plain text and HTML. In future iterations, we would like to add support for more reading formats, as well as support to interface with other online book service providers.  

## Technical Requirements 

- Get list of books
- Get a book
- Get a book page in the desired format
- Make use of friendly routes (for example: `/book/1` or `/book/1 /page/11/html`)
- Provide seeders / migrations for the database (books with their pages)
- Provide instruccions on project setup/configuration

## Rules
- Do not use any frameworks. A good developer must know how to select his tools and how to use them.
- The use of third-party libraries is allowed and encouraged.

## Evaluation Criteria

1. Technical requirements.
2. Organization and consistency of the file and folder structure.
3. Modifiability and extendability of the system where required. 
4. Commit history (commits are organized and descriptive).
5. Time used to complete the test.
6. Complexity of the solution.
7. Correct usage of SOLID principles.
8. Correct usage of design patterns.

## Let's Go

### Prerequisites

- An installation of the MySQL relational database management system (DBMS) https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/.
- An installation of Go. For installation instructions, see Installing Go https://go.dev/doc/install.
- A tool to edit your code. Any text editor you have will work fine.
- A command terminal. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

### Set up a database

We’ll use the CLI for the DBMS itself to create the database and table, as well as to add data.
The code here uses the MySQL CLI
1. Open a new command prompt.
2. Export mysql to env
```bash
export PATH=$PATH:/usr/local/mysql/bin
```
3. At the command line, log into your DBMS, as in the following example for MySQL.
```bash
mysql -u root -p
Enter password:
```
4. At the mysql command prompt, create a database.
```bash
create database {DB_NAME};
```
5. Change to the database you just created so you can add tables.
```bash
use {DB_NAME};
```
6. From the mysql command prompt, run the script in create-tables.sql file.
```bash
source /path/to/create-tables.sql
```
7. At your DBMS command prompt, use a SELECT statement to verify you’ve successfully created the table with data.
```bash
select * from books;
```

### Getting Started

1. Clone the repo from main branch into your local machine.
2. From terminal execute command"
```bash
cd book-lib-api
```
3. From terminal, add project dependencies:
```bash
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
go get github.com/DATA-DOG/go-sqlmock
```
4. Try to find any missing dependencies
```bash
go mod tidy
```
5. Fill env variables in .env file 
6. Run the program
```bash
go run .
```
