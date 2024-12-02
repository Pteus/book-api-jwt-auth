package services

import (
	"github.com/pteus/books-api/internal/models"
	"github.com/pteus/books-api/internal/repositories"
)

type BookService interface {
	Create(book *models.Book) error
	GetAllByUsername(username string) ([]models.Book, error)
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{bookRepo}
}

func (b *bookService) Create(book *models.Book) error {
	err := b.bookRepo.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookService) GetAllByUsername(username string) ([]models.Book, error) {
	books, err := b.bookRepo.GetAllByUsername(username)
	if err != nil {
		return nil, err
	}

	return books, nil
}
