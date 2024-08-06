package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMember(member models.Member)
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}

func (l *Library) AddMember(member models.Member) {
	l.members[member.ID] = member
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Available"
	l.books[bookID] = book

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
