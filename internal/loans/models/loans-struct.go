package models

import (
	"time"
	"github.com/google/uuid"
)

type Loan struct {
	Id         uuid.UUID     `json:"id"`
	BookId     uuid.UUID     `json:"bookId"`
	UserId     uuid.UUID     `json:"userId"`
	BorrowedAt time.Time 		 `json:"borrowedAt"`
	ReturnedAt time.Time 		 `json:"returnedAt"`
	Status     string    		 `json:"status"`
	CreatedAt  time.Time 		 `json:"createdAt"`
	UpdatedAt  time.Time 		 `json:"updatedAt"`
}