package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	library services.LibraryManager
}

func NewLibraryController(library services.LibraryManager) *LibraryController {
	return &LibraryController{library: library}
}

func (lc *LibraryController) AddBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	lc.library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) RemoveBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	err := lc.library.RemoveBook(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func (lc *LibraryController) BorrowBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, _ := strconv.Atoi(bookIDStr)

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, _ := strconv.Atoi(memberIDStr)

	err := lc.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (lc *LibraryController) ReturnBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, _ := strconv.Atoi(bookIDStr)

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, _ := strconv.Atoi(memberIDStr)

	err := lc.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
	} else {
		fmt.Println("Available Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func (lc *LibraryController) ListBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberIDStr = strings.TrimSpace(memberIDStr)
	memberID, _ := strconv.Atoi(memberIDStr)

	books := lc.library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books.")
	} else {
		fmt.Println("Borrowed Books:")
		for _, book := range books {
			fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func (lc *LibraryController) AddMember(reader *bufio.Reader) {
	fmt.Print("Enter member ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	fmt.Print("Enter member name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	member := models.Member{
		ID:   id,
		Name: name,
	}
	lc.library.AddMember(member)
	fmt.Println("Member added successfully!")
}