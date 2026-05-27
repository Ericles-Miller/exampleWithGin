package repositories

import (
	"exampleWithGin/internal/users/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) models.UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return nil
}

func (r *UserRepository) GetUser(id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	return nil, nil
}

func (r *UserRepository) UpdateUser(id uuid.UUID, user *models.User) error {
	return nil
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	return nil
}
