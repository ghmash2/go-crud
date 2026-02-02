package service

import (
	"context"
	"go-crud/internal/models"
	"go-crud/internal/repository"
	"go-crud/internal/validation"
	"strconv"

	"github.com/shopspring/decimal"
)

type ProductService interface {
	CreateProduct(ctx context.Context, in validation.ProductValidation) (uint64, error)
	GetProductByID(ctx context.Context, id uint64) (models.Product, error)
	UpdateProduct(ctx context.Context, id uint64, in validation.ProductValidation) (uint64, error)
	DeleteProduct(ctx context.Context, id uint64) error
	GetAllProducts(ctx context.Context) ([]models.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s productService) CreateProduct(ctx context.Context, in validation.ProductValidation) (uint64, error) {
	if err := validation.ValidateProduct(in); err != nil {
		return 0, err
	}

	// Convert price string to decimal.Decimal
	price, err := decimal.NewFromString(in.Price)
	if err != nil {
		return 0, err
	}

	// Convert categoryID string to uint64
	categoryID, err := strconv.ParseUint(in.CategoryID, 10, 64)
	if err != nil {
		return 0, err
	}

	product := models.Product{
		Name:       in.Name,
		Price:      price,
		CategoryID: categoryID,
	}
	return s.productRepo.CreateProduct(ctx, product)
}

func (s productService) GetProductByID(ctx context.Context, id uint64) (models.Product, error) {
	return s.productRepo.GetProductByID(ctx, id)
}

func (s productService) UpdateProduct(ctx context.Context, id uint64, in validation.ProductValidation) (uint64, error) {
	if err := validation.ValidateProduct(in); err != nil {
		return 0, err
	}

	// Convert price string to decimal.Decimal
	price, err := decimal.NewFromString(in.Price)
	if err != nil {
		return 0, err
	}

	// Convert categoryID string to uint64
	categoryID, err := strconv.ParseUint(in.CategoryID, 10, 64)
	if err != nil {
		return 0, err
	}

	product := models.Product{
		Name:       in.Name,
		Price:      price,
		CategoryID: categoryID,
	}
	return s.productRepo.UpdateProduct(ctx, id, product)
}

func (s productService) DeleteProduct(ctx context.Context, id uint64) error {
	return s.productRepo.DeleteProduct(ctx, id)
}

func (s productService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.productRepo.GetAllProducts(ctx)
}
