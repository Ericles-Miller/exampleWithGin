package models

import "github.com/google/uuid"

type UserService interface {
	CreateUser(user *User) error
	GetUser(id uuid.UUID) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(id uuid.UUID, user *User) (*User, error)
	DeleteUser(id uuid.UUID) error
}