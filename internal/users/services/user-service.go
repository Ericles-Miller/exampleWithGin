package services

import (
	"exampleWithGin/internal/users/models"
	"exampleWithGin/pkg/appErrors"
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
	userAlreadyExists, _ := u.userRepo.GetUserByEmail(user.Email)
	if userAlreadyExists != nil {
		return nil, appErrors.ErrBadRequest
	}

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) GetUser(id uuid.UUID) (*models.User, error) {
	user, err := u.userRepo.GetUser(id)
	if err != nil {
		return nil, appErrors.ErrNotFound
	}
	return user, nil
}

func (u UserService) GetAllUsers() ([]*models.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u UserService) UpdateUser(id uuid.UUID, user *models.User) (*models.User, error) {
	user, err := u.GetUser(id)
	if err != nil {
		return nil, appErrors.ErrNotFound
	}
	
	user.UpdatedAt = time.Now()
	err = u.userRepo.UpdateUser(id, user)

	return user, nil
}

func (u UserService) DeleteUser(id uuid.UUID) error {
	user, err := u.GetUser(id)
	if err != nil {
		return appErrors.ErrNotFound
	}
	
	err = u.userRepo.DeleteUser(user.Id)

	return nil
}
