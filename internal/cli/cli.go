package cli

import (
	"bufio"
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"strings"
	"time"
)

type BookCLI struct {
	bookService *service.BookService
}

func NewBookCLI(bookService *service.BookService) *BookCLI {
	return &BookCLI{bookService: bookService}
}

func (cli *BookCLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("<command> <book name> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		command := args[0]

		switch command {
		case "search":
			if len(args) < 2 {
				fmt.Println("Usage: search <book title>")
				continue
			}
			bookName := strings.Join(args[1:], " ")
			cli.searchBooks(bookName)

		case "simulate":
			if len(args) < 2 {
				fmt.Println("Usage: simulate <book_id> <book_id> ...")
				continue
			}
			bookIDs := args[1:]
			cli.simulateReading(bookIDs)

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown command:", command)
		}
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
