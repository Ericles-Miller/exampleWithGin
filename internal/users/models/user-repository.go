package models

import "github.com/google/uuid"

type UserRepository interface {
	CreateUser(user *User) error
	GetUser(id uuid.UUID) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(id uuid.UUID, user *User) error
	DeleteUser(id uuid.UUID) error
	GetUserByEmail(email string) (*User, error)
}