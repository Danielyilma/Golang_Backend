# Library Management System (Go)

## Project Overview

This project is a **console-based Library Management System** implemented in Go. It follows object-oriented principles using interfaces and struct composition. The goal is to simulate a simple library where users can manage books and members.

## Key Features

* Add and remove books from the library
* Register members and track borrowed books
* Borrow and return books
* Display available books
* Display books borrowed by a specific member

## Folder Structure

* `main.go` - Entry point of the application; handles the user interface and CLI menu.
* `controllers/` - Contains controller functions that handle user interaction logic.
* `services/` - Business logic for managing books and members (implementing interfaces).
* `models/` - Data structures (e.g., `Book`, `Member`) used throughout the application.

## Sample Struct Definitions

```go
/**
 * struct Member - Struct representing a library member
 * @Id: Unique identifier for the member
 * @Name: Name of the member
 * @BorowwedBooks: List of books currently borrowed by the member
 */
type Member struct {
	Id             int
	Name           string
	BorowwedBooks  []Book
}

/**
 * struct Book - Struct representing a book in the library
 * @Id: Unique identifier for the book
 * @Title: Title of the book
 * @Author: Author of the book
 * @Status: Status of the book ("Available" or "Not Available")
 */
type Book struct {
	Id     int
	Title  string
	Author string
	Status string
}
```

## Usage

Run the program using:

```bash
go run main.go
```

Follow the CLI menu to manage books and members.

## Design Patterns Used

* Singleton pattern for the library instance
* Interface-based abstraction for library management

## Future Improvements

* Persistent storage (e.g., save to file or database)
* Advanced search and filtering
* Web or GUI interface
