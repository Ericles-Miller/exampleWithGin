package models

import "github.com/google/uuid"

type LoansService interface {
	CreateLoan(loan *Loan) (*Loan, error)
	GetLoan(id uuid.UUID) (*Loan, error)
	GetAllLoans() ([]*Loan, error)
	UpdateLoan(id uuid.UUID, loan *Loan) (*Loan, error)
	DeleteLoan(id uuid.UUID) error
}