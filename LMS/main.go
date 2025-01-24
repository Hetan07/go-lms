package main

import (
	"LMS/book"
	"LMS/library"
	"fmt"
)

func main() {
	lib := &library.Library{}

	book1 := book.NewBook("The Go Programming Language", "Alan A. A. Donovan & Brian W. Kernighan", "978-0134190440", true)
	book2 := book.NewBook("Clean Code", "Robert C. Martin", "978-0132350884", false)
	lib.AddBook(book1)
	lib.AddBook(book2)

	fmt.Println("All Books in the Library:")
	lib.ListBooks()

	fmt.Println("\nSearch Results for 'Go':")
	results := lib.SearchBookByTitle("Go")
	for _, b := range results {
		fmt.Println(b.Display())
	}

	fmt.Println("\nRemoving book with ISBN 978-0132350884...")
	lib.RemoveBook("978-0132350884")

	fmt.Println("\nAll Books in the Library After Removal:")
	lib.ListBooks()
}
