package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validate = validator.New()

type User struct {
	Id        uuid.UUID `json:"id"        validate:"required,uuid4"`
	Name      string    `json:"name"      validate:"required,min=10,max=200"`
	Email     string    `json:"email"     validate:"required,email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
