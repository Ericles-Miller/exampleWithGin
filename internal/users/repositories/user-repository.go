package repositories

import (
	"context"
	"exampleWithGin/internal/users/db"
	"exampleWithGin/internal/users/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	queries *userdb.Queries
}

func NewUserRepository(pool *pgxpool.Pool) models.UserRepository {
	return &UserRepository{
		queries: userdb.New(pool),
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	result, err := r.queries.CreateUser(context.Background(), userdb.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return err
	}

	user.Id = result.ID
	user.CreatedAt = result.CreatedAt.Time
	user.UpdatedAt = result.UpdatedAt.Time
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	result, err := r.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetUser(id uuid.UUID) (*models.User, error) {
	result, err := r.queries.GetUser(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	results, err := r.queries.GetAllUsers(context.Background())
	if err != nil {
		return nil, err
	}

	users := make([]*models.User, len(results))
	for i, result := range results {
		users[i] = &models.User{
			Id:        result.ID,
			Name:      result.Name,
			Email:     result.Email,
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.UpdatedAt.Time,
		}
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(id uuid.UUID, user *models.User) error {
	result, err := r.queries.UpdateUser(context.Background(), userdb.UpdateUserParams{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return err
	}

	user.UpdatedAt = result.UpdatedAt.Time
	return nil
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	return r.queries.DeleteUser(context.Background(), id)
}
