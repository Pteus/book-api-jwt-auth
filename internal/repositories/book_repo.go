package repositories

import (
	"database/sql"

	"github.com/pteus/books-api/internal/models"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetAllByUsername(username string) ([]models.Book, error)
	GetByID(id int, username string) (*models.Book, error)
	Update(id int, book *models.Book, username string) error
	Delete(id int, username string) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db}
}

func (b *bookRepository) CreateBook(book *models.Book) error {
	_, err := b.db.Exec("INSERT INTO books (title, author, genre, username) values (_,err,err,)", book.Title, book.Author, book.Genre, book.Username)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookRepository) GetAllByUsername(username string) ([]models.Book, error) {
	rows, err := b.db.Query("SELECT id, title, author, genre WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *bookRepository) GetByID(id int, username string) (*models.Book, error) {
	panic("not implemented") // TODO: Implement
}

func (b *bookRepository) Update(id int, book *models.Book, username string) error {
	panic("not implemented") // TODO: Implement
}

func (b *bookRepository) Delete(id int, username string) error {
	panic("not implemented") // TODO: Implement
}
