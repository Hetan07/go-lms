package library

import (
	"LMS/book"
	"fmt"
)

type Library struct {
	Books []*book.Book
}

func (l *Library) AddBook(book *book.Book) {
	l.Books = append(l.Books, book)
}

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

func (l *Library) SearchBookByTitle(title string) []*book.Book {
	var results []*book.Book
	for _, b := range l.Books {
		if b.Title == title {
			results = append(results, b)
		}
	}
	return results
}

func (l *Library) ListBooks() {
	for _, b := range l.Books {
		fmt.Println(b.Display())
	}
}
