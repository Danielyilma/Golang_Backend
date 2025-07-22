package controllers

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/Danielyilma/library_management/models"
	"github.com/Danielyilma/library_management/services"
)

var libInstance *services.Library
var id_num int = 0


/**
 * NewLibrary - Initializes a new Library instance if not already created.
 * Return: LibraryManager interface to interact with the service layer.
 */

func NewLibrary() services.LibraryManager {
	if libInstance == nil {
		libInstance = &services.Library{
			Books:   make(map[int]models.Book),
			Members: make(map[int]models.Member),
		}
	}
	return libInstance
}


/**
 * AddBook - Handles input and adds a new book to the library.
 */

func AddBook() {
	var book models.Book
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\t\tEnter Book Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("\t\tError reading title:", err)
		return
	}
	book.Title = strings.TrimSpace(title)

	fmt.Print("\t\tEnter Book Author: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("\t\tError reading author:", err)
		return
	}
	book.Author = strings.TrimSpace(author)

	book.Id = id_num
	book.Status = "Available"
	id_num++

	libInstance.AddBook(book)
	fmt.Println("\t\tBook added successfully!")
}


/**
 * RemoveBook - Prompts for a book ID and removes the book if found.
 */

func RemoveBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\t\tEnter Book ID to remove: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("\t\tError reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("\t\tInvalid ID format")
		return
	}

	err = libInstance.RemoveBook(id)
	if err != nil {
		fmt.Println("\t\tError:", err)
		return
	}

	fmt.Println("\t\tBook removed successfully!")
}


/**
 * BorrowBookController - Reads member and book IDs and attempts to borrow the book.
 * @service: LibraryManager interface instance
 */

func BorrowBookController(service services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\t\tEnter Member ID: ")
	memberInput, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(memberInput))
	if err != nil {
		fmt.Println("\t\tInvalid Member ID")
		return
	}

	fmt.Print("\t\tEnter Book ID to borrow: ")
	bookInput, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(bookInput))
	if err != nil {
		fmt.Println("\t\tInvalid Book ID")
		return
	}

	err = libInstance.BorrowBook(memberID, bookID)
	if err != nil {
		fmt.Println("\t\tError:", err)
		return
	}
	fmt.Println("\t\tBook borrowed successfully!")
}


/**
 * ReturnBookController - Reads member and book IDs and returns the borrowed book.
 * @service: LibraryManager interface instance
 */

func ReturnBookController(service services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\t\tEnter Member ID: ")
	memberInput, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(memberInput))
	if err != nil {
		fmt.Println("\t\tInvalid Member ID")
		return
	}

	fmt.Print("\t\tEnter Book ID to return: ")
	bookInput, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(bookInput))
	if err != nil {
		fmt.Println("\t\tInvalid Book ID")
		return
	}

	err = libInstance.ReturnBook(memberID, bookID)
	if err != nil {
		fmt.Println("\t\tError:", err)
		return
	}
	fmt.Println("\t\tBook returned successfully!")
}


/**
 * ListAvailableBooksController - Displays all available books in the library.
 * @service: LibraryManager interface instance
 */

func ListAvailableBooksController(service services.LibraryManager) {
	books := libInstance.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("\t\tNo books are currently available.")
		return
	}

	fmt.Println("\t\tAvailable Books:")
	for _, book := range books {
		fmt.Printf("\t\tID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}


/**
 * ListBorrowedBooksByMemberController - Lists books borrowed by a specific member.
 * @service: LibraryManager interface instance
 */

func ListBorrowedBooksByMemberController(service services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\t\tEnter Member ID: ")
	memberInput, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(memberInput))
	if err != nil {
		fmt.Println("\t\tInvalid Member ID")
		return
	}

	books := libInstance.ListBorrowedBooks(memberID)

	if len(books) == 0 {
		fmt.Println("\t\tNo books borrowed by this member.")
		return
	}

	fmt.Println("\t\tBorrowed Books:")
	for _, book := range books {
		fmt.Printf("\t\tID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}
