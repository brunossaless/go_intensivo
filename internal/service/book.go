package service

import (
	"database/sql"
	"log"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := "INSERT INTO books (title, author, genre) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		log.Printf("Erro ao executar query: %v\n", err)
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Erro ao obter LastInsertId: %v\n", err)
		return err
	}

	if lastInsertId == 0 {
		log.Println("Nenhum ID foi retornado, verifique a configuração da coluna ID")
	}

	book.ID = int(lastInsertId)
	log.Printf("Livro criado com ID: %d\n", book.ID)
	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := "Select id, title, author, genre from books"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*Book, error) {
	query := "Select id, title, author, genre from books where id = ?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := "update book set title=?, author=?, genre=? where id=?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)

	return err
}

func (s *BookService) DeleteBook(id int) error {
	query := "delete from books where id=?"
	_, err := s.db.Exec(query, id)

	return err
}
