package models

import "github.com/google/uuid"

type BookRepository interface {
	CreateBook(book *Book) (*Book, error)
	GetBook(id uuid.UUID) (*Book, error)
	GetAllBooks() ([]*Book, error)
	UpdateBook(id uuid.UUID, book *Book) (*Book,error)
	DeleteBook(id uuid.UUID) error
}