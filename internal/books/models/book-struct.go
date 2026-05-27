package models

import "github.com/google/uuid"

type Book struct {
	Id     uuid.UUID `json:"id" validate:"required, uuid4"`
	Title  string `json:"title" validate:"required, min=3, max=100"`
	Author string `json:"author" validate:"required,min=3,max=100"`
	Quantity  int `json:"quantity" validate:"required, min=0"`
}