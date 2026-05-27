package services

import (
	"errors"
	bookModels "exampleWithGin/internal/books/models"
	"exampleWithGin/internal/loans/models"
	userModels "exampleWithGin/internal/users/models"
	"time"

	"github.com/google/uuid"
)

type LoansService struct {
	loanRepository models.LoanRepository
	bookService    bookModels.BookService
	userService    userModels.UserService
}

func NewLoanService(
	loanRepository models.LoanRepository,
	bookService bookModels.BookService,
	userService userModels.UserService,
) models.LoansService {
	return &LoansService{
		loanRepository: loanRepository,
		bookService:    bookService,
		userService:    userService,
	}
}

func (l LoansService) CreateLoan(loan *models.Loan) (*models.Loan, error) {
	book, err := l.bookService.GetBook(loan.BookId)
	if err != nil {
		return nil, err
	}

	if book.Quantity <= 0 {
		return nil, errors.New("book is not available")
	}

	_, err = l.userService.GetUser(loan.UserId)
	if err != nil {
		return nil, err
	}

	activeLoans, err := l.loanRepository.GetActiveUserLoans(loan.UserId)
	if err != nil {
		return nil, err
	}

	if len(activeLoans) > 0 {
		return nil, errors.New("user has active loans")
	}

	loan.BorrowedAt = time.Now()
	loan.Status = "active"
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()

	loan, err = l.loanRepository.CreateLoan(loan)
	if err != nil {
		return nil, err
	}

	book.Quantity--
	if _, err = l.bookService.UpdateBook(book.Id, book); err != nil {
		return nil, err
	}

	return loan, nil
}

func (l LoansService) ReturnBook(loanId uuid.UUID) error {
	loan, err := l.loanRepository.GetLoan(loanId)
	if err != nil {
		return err
	}

	if loan.Status == "returned" {
		return errors.New("book already returned")
	}

	loan.Status = "returned"
	loan.UpdatedAt = time.Now()
	loan.ReturnedAt = time.Now()

	if err := l.loanRepository.ReturnBook(loan); err != nil {
		return err
	}

	book, err := l.bookService.GetBook(loan.BookId)
	if err != nil {
		return err
	}

	book.Quantity++
	_, err = l.bookService.UpdateBook(book.Id, book)
	return err
}

func (l LoansService) GetLoan(id uuid.UUID) (*models.Loan, error) {
	return l.loanRepository.GetLoan(id)
}

func (l LoansService) GetAllLoans() ([]*models.Loan, error) {
	return l.loanRepository.GetAllLoans()
}

func (l LoansService) UpdateLoan(id uuid.UUID, loan *models.Loan) (*models.Loan, error) {
	loan.Id = id
	loan.UpdatedAt = time.Now()
	return l.loanRepository.CreateLoan(loan)
}

func (l LoansService) DeleteLoan(id uuid.UUID) error {
	return errors.New("not implemented")
}
