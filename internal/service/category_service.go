package service

import (
	"context"
	"go-crud/internal/models"
	"go-crud/internal/repository"
	"go-crud/internal/validation"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, in validation.CategoryValidation) (uint64, error)
	GetCategoryByID(ctx context.Context, id uint64) (models.Category, error)
	UpdateCategory(ctx context.Context, id uint64, in validation.CategoryValidation) (uint64, error)
	DeleteCategory(ctx context.Context, id uint64) error
	GetAllCategories(ctx context.Context) ([]models.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s categoryService) CreateCategory(ctx context.Context, in validation.CategoryValidation) (uint64, error) {
	if err := validation.ValidateCategory(in); err != nil {
		return 0, err
	}
	category := models.Category{
		Name: in.Name,
	}
	return s.categoryRepo.CreateCategory(ctx, category)
}

func (s categoryService) GetCategoryByID(ctx context.Context, id uint64) (models.Category, error) {
	return s.categoryRepo.GetCategoryByID(ctx, id)
}

func (s categoryService) UpdateCategory(ctx context.Context, id uint64, in validation.CategoryValidation) (uint64, error) {
	if err := validation.ValidateCategory(in); err != nil {
		return 0, err
	}
	category := models.Category{
		Name: in.Name,
	}
	return s.categoryRepo.UpdateCategory(ctx, id, category)
}

func (s categoryService) DeleteCategory(ctx context.Context, id uint64) error {
	return s.categoryRepo.DeleteCategory(ctx, id)
}

func (s categoryService) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return s.categoryRepo.GetAllCategories(ctx)
}