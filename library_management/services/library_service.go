package services

import (
	"fmt"
	"github.com/Danielyilma/library_management/models"
)


/**
 * LibraryManager - Interface defining core library operations
 */

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

/**
 * Library - Struct to manage books and members in the library
 * @Books: map of all books by their ID
 * @Members: map of all members by their ID
 */

type Library struct {
	Books  map[int] models.Book
	Members map[int] models.Member
}

/**
 * checkBookExists - Checks if a book with given ID exists
 * @bookID: ID of the book to check
 * Return: pointer to the book or an error
 */

func (L *Library) checkBookExists(bookID int) (*models.Book, error) {
	book, ok := L.Books[bookID]

	if !ok {
		return nil, fmt.Errorf("book not found with this id %d", bookID)
	}

	return &book, nil
}

/**
 * checkMemberExists - Checks if a member with given ID exists
 * @memberID: ID of the member to check
 * Return: pointer to the member or an error
 */

func (L *Library) checkMemberExists(memberID int) (*models.Member, error) {
	member, ok := L.Members[memberID]

	if !ok {
		return nil, fmt.Errorf("member not found with this id %d", memberID)
	}

	return &member, nil
}

/**
 * AddBook - Adds a new book to the library
 * @book: Book object to add
 */

func (L *Library) AddBook(book models.Book) {
	L.Books[book.Id] = book
}

/**
 * RemoveBook - Removes a book from the library by ID
 * @bookID: ID of the book to remove
 * Return: error if book does not exist
 */

func (L * Library) RemoveBook(bookID int) error {
	if _, err := L.checkBookExists(bookID); err != nil {
		return err
	}

	delete(L.Books, bookID)
	return nil
}

/**
 * BorrowBook - Allows a member to borrow a book
 * @bookID: ID of the book to borrow
 * @memberID: ID of the member borrowing the book
 * Return: error if book or member not found
 */

func (L * Library) BorrowBook(bookID,  memberID int) error {
	book, err := L.checkBookExists(bookID)
	if err != nil {
		return err
	}

	member, err := L.checkMemberExists(memberID)
	if err != nil {
		return err
	}

	member.BorowwedBooks = append(member.BorowwedBooks, *book)
	book.Status = "Not Available"
	
	return nil
}

/**
 * ReturnBook - Allows a member to return a borrowed book
 * @bookID: ID of the book to return
 * @memberID: ID of the member returning the book
 * Return: error if book was not borrowed
 */

func (L * Library) ReturnBook(bookID int, memberID int) error {
	book, err := L.checkBookExists(bookID)
	if err != nil {
		return err
	}

	member, err := L.checkMemberExists(memberID)
	if err != nil {
		return err
	}

	for i, b := range member.BorowwedBooks {
		if b.Id == bookID {
			member.BorowwedBooks[i] = member.BorowwedBooks[len(member.BorowwedBooks)-1]
			member.BorowwedBooks = member.BorowwedBooks[:len(member.BorowwedBooks)-1]
			book.Status = "Available"
			return nil
		}
	}
	
	return fmt.Errorf("member %d did not borrow book %d", memberID, bookID)
}

/**
 * ListAvailableBooks - Lists all books currently in the library
 * Return: slice of all books in the library
 */

func (L * Library) ListAvailableBooks() []models.Book {
	var books []models.Book

	for _, book := range L.Books {
		books = append(books, book)
	}

	return books
}

/**
 * ListBorrowedBooks - Lists all books borrowed by a specific member
 * @memberID: ID of the member
 * Return: slice of borrowed books or empty if member doesn't exist
 */

func (L * Library) ListBorrowedBooks(memberID int) []models.Book{
	member, err := L.checkMemberExists(memberID)
	if err != nil {
		return []models.Book{}
	}

	return member.BorowwedBooks
}

