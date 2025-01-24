package book

import (
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func NewBook(title string, author string, isbn string, avail bool) *Book {
	return &Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: avail,
	}
}

func (b *Book) Display() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nStatus: %t\n", b.Title, b.Author, b.ISBN, b.Available)
}

func main() {
	book := NewBook("1", "2", "3", true)
	fmt.Println(book.Display())
}
