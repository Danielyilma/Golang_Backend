package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Danielyilma/library_management/controllers"
)

/*
 * main - Entry point of the Library Management System
 * Description: Displays a menu for managing books and members.
 */


func main() {
	reader := bufio.NewReader(os.Stdin)
	lib_service := controllers.NewLibrary()
	fmt.Println("\n\t\tWelcome to the Library Management System")

	for {
		fmt.Println("\n\t\t\tMenu:")
		fmt.Println("\t\t\t1. Add Book")
		fmt.Println("\t\t\t2. Remove Book")
		fmt.Println("\t\t\t3. Borrow Book")
		fmt.Println("\t\t\t4. Return Book")
		fmt.Println("\t\t\t5. List Available Books")
		fmt.Println("\t\t\t6. List Borrowed Books by Member")
		fmt.Println("\t\t\t7. Exit\n\n")

		fmt.Print("\t\tChoose an option: ")
		input, _ := reader.ReadString('\n')
		fmt.Println()
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			controllers.AddBook()

		case "2":
			controllers.RemoveBook()

		case "3":
			controllers.BorrowBookController(lib_service)

		case "4":
			controllers.ReturnBookController(lib_service)

		case "5":
			controllers.ListAvailableBooksController(lib_service)

		case "6":
			controllers.ListBorrowedBooksByMemberController(lib_service)

		case "7":
			fmt.Println("\t\tGoodbye!")
			return

		default:
			fmt.Println("\t\tInvalid option, please try again.")
		}

		fmt.Println("\n\n\n")
	}
}
