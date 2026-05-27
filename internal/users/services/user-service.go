package services

import (
	"exampleWithGin/internal/users/models"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) models.UserService {
	return &UserService{userRepo: userRepo}
}

func (u UserService) CreateUser(user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) GetUser(id uuid.UUID) (*models.User, error) {
	return u.userRepo.GetUser(id)
}

func (u UserService) GetAllUsers() ([]*models.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u UserService) UpdateUser(id uuid.UUID, user *models.User) (*models.User, error) {
	user.UpdatedAt = time.Now()
	err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) DeleteUser(id uuid.UUID) error {
	return u.userRepo.DeleteUser(id)
}
