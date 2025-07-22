package models

/**
 * struct Member - Struct representing a library member
 * @Id: Unique identifier for the member
 * @Name: Name of the member
 * @BorowwedBooks: List of books currently borrowed by the member
 */

type Member struct {
	Id 	int
	Name string
	BorowwedBooks []Book
}