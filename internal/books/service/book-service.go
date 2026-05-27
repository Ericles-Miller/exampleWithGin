package services

import (
	"errors"
	"exampleWithGin/internal/books/models"
	"github.com/google/uuid"
)

type BookService struct {
	bookRepository models.BookRepository
}

func NewBookService(bookRepository models.BookRepository) models.BookService {
	return &BookService{bookRepository: bookRepository}
}

func (b BookService) CreateBook(book *models.Book) (*models.Book, error) {
	if book.Title == "" {
		return nil, errors.New("title is required")
	}
	if book.Author == "" {
		return nil, errors.New("author is required")
	}
	if book.Quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}
	return b.bookRepository.CreateBook(book)
}

func (b BookService) GetBook(id uuid.UUID) (*models.Book, error) {
	return b.bookRepository.GetBook(id)
}

func (b BookService) GetAllBooks() ([]*models.Book, error) {
	return b.bookRepository.GetAllBooks()
}

func (b BookService) UpdateBook(id uuid.UUID, book *models.Book) (*models.Book, error) {
	return b.bookRepository.UpdateBook(id, book)
}

func (b BookService) DeleteBook(id uuid.UUID) error {
	return b.bookRepository.DeleteBook(id) 
}
