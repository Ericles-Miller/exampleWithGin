package models

type Loan struct {
	ID		int	`json:"ID"`
	UserID	int	`json:"userId"`
	BookID	int	`json:"bookId"`
}