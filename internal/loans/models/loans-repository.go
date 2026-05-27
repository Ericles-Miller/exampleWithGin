package models

import "github.com/google/uuid"

type LoanRepository interface {
	CreateLoan(loan *Loan) (*Loan, error)
	ReturnBook(loan *Loan) error
	GetLoan(id uuid.UUID) (*Loan, error)
	GetActiveUserLoans(userId uuid.UUID) ([]*Loan, error)
	GetAllLoans() ([]*Loan, error)
}