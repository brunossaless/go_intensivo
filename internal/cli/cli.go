package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	bookService *service.BookService
}

func NewBookCLI(bookService *service.BookService) *BookCLI {
	return &BookCLI{bookService: bookService}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: books <command> [arguments...]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)

	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> ...")
			return
		}

		//Pega todos do 2 args pra frente
		bookIDs := os.Args[2:]
		cli.simulateReading(bookIDs)
	}

}

func (cli *BookCLI) searchBooks(bookName string) {
	books, err := cli.bookService.SearchBooksByName(bookName)
	if err != nil {
		fmt.Printf("Error searching books: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Printf("No books found with title '%s'\n", bookName)
		return
	}

	fmt.Printf("%d Found books:", len(books))
	for _, book := range books {
		fmt.Printf("\nID: %d\nTitle: %s\nAuthor: %s\nGenre: %s",
			book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReading(bookIDsStr []string) {
	var bookIDs []int
	for _, idString := range bookIDsStr {
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Printf("Invalid book ID: %s\n", idString)
			continue
		}

		bookIDs = append(bookIDs, id)
	}

	responses := cli.bookService.SimulateMultipleRead(bookIDs, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}
