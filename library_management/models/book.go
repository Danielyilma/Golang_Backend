package models

/**
 * struct Book - Struct representing a book in the library
 * @Id: Unique identifier for the book
 * @Title: Title of the book
 * @Author: Author of the book
 * @Status: Status of the book ("Available" or "Not Available")
 */

type Book struct {
	Id	 int
	Title  string
	Author string
	Status string
}
