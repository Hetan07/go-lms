package main

import (
	"fmt"
	"LMS/Bo"    // Replace "your-module-name" with your actual module name
	"your-module-name/library" // Replace "your-module-name" with your actual module name
)

func main() {
	// Create a new library
	lib := &library.Library{}

	// Add some books
	book1 := book.NewBook("The Go Programming Language", "Alan A. A. Donovan & Brian W. Kernighan", "978-0134190440", true)
	book2 := book.NewBook("Clean Code", "Robert C. Martin", "978-0132350884", false)
	lib.AddBook(book1)
	lib.AddBook(book2)

	// List all books
	fmt.Println("All Books in the Library:")
	lib.ListBooks()

	// Search for books by title
	fmt.Println("\nSearch Results for 'Go':")
	results := lib.SearchBookByTitle("Go")
	for _, b := range results {
		fmt.Println(b.Display())
	}

	// Remove a book by ISBN
	fmt.Println("\nRemoving book with ISBN 978-0132350884...")
	lib.RemoveBook("978-0132350884")

	// List all books after removal
	fmt.Println("\nAll Books in the Library After Removal:")
	lib.ListBooks()
}