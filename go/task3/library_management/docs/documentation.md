# Library Management System

## Overview
This is a simple console-based library management system implemented in Go. It demonstrates the use of structs, interfaces, methods, slices, and maps.

## Features
- Add a new book to the library
- Add a new member to the library
- Remove an existing book from the library
- Borrow a book
- Return a book
- List all available books
- List all borrowed books by a member

## Usage
Run the `main.go` file to start the application. The system will prompt you to choose from various options to interact with the library.

## Structs
### Book
- ID (int)
- Title (string)
- Author (string)
- Status (string) // "Available" or "Borrowed"

### Member
- ID (int)
- Name (string)
- BorrowedBooks ([]Book) // a slice to hold borrowed books

## Interfaces
### LibraryManager
- AddBook(book Book)
- AddMember(memeber Member)
- RemoveBook(bookID int)
- BorrowBook(bookID int, memberID int) error
- ReturnBook(bookID int, memberID int) error
- ListAvailableBooks() []Book
- ListBorrowedBooks(memberID int) []Book

## Implementation
The `Library` struct implements the `LibraryManager` interface and contains the following methods:
- AddBook
- AddMember
- RemoveBook
- BorrowBook
- ReturnBook
- ListAvailableBooks
- ListBorrowedBooks

## Folders
- `controllers/`: Handles console input and invokes the appropriate service methods.
- `models/`: Defines the `Book` and `Member` structs.
- `services/`: Contains business logic and data manipulation functions.
- `docs/`: Contains system documentation and other related information.
