package service

import (
	"context"
	"go-crud/internal/models"
	"go-crud/internal/repository"
	"go-crud/internal/validation"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	CreateUser(ctx context.Context, in validation.UserValidation) (uint64, error)
	GetUserByID(ctx context.Context, id uint64) (models.User, error)
	UpdateUser(ctx context.Context, id uint64, in validation.UserValidation) (uint64, error)
	DeleteUser(ctx context.Context, id uint64) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s userService) CreateUser(ctx context.Context, in validation.UserValidation) (uint64, error) {

	if err := validation.ValidateUser(in); err != nil {
		return 0, err
	}
	user := models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	return s.userRepo.CreateUser(ctx, user)
}

func (s userService) GetUserByID(ctx context.Context, id uint64) (models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s userService) UpdateUser(ctx context.Context, id uint64, in validation.UserValidation) (uint64, error) {
	if err := validation.ValidateUser(in); err != nil {
		return 0, err
	}
	user := models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	return s.userRepo.UpdateUser(ctx, id, user)
}

func (s userService) DeleteUser(ctx context.Context, id uint64) error {
	return s.userRepo.DeleteUser(ctx, id)
}

func (s userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}
