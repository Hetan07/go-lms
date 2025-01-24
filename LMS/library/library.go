package library

import (
	"fmt"
	"strings"
	"your-module-name/book" // Replace "your-module-name" with your actual module name
)

// Define the Library struct
type Library struct {
	Books []*book.Book
}

// Method to add a book to the library
func (l *Library) AddBook(book *book.Book) {
	l.Books = append(l.Books, book)
}

// Method to remove a book from the library using its ISBN
func (l *Library) RemoveBook(isbn string) {
	for i, b := range l.Books {
		if b.ISBN == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Printf("Book with ISBN %s removed.\n", isbn)
			return
		}
	}
	fmt.Printf("Book with ISBN %s not found.\n", isbn)
}

// Method to search for books by title (partial or full match)
func (l *Library) SearchBookByTitle(title string) []*book.Book {
	var results []*book.Book
	for _, b := range l.Books {
		if B.title == title {
			results = append(results, b)
		}
	}
	return results
}

// Method to list all books in the library
func (l *Library) ListBooks() {
	for _, b := range l.Books {
		fmt.Println(b.Display())
	}
}