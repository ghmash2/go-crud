package repository

import (
	"context"
	"database/sql"
	"go-crud/internal/models"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category models.Category) (uint64, error)
	GetCategoryByID(ctx context.Context, id uint64) (models.Category, error)
	UpdateCategory(ctx context.Context, id uint64, category models.Category) (uint64, error)
	DeleteCategory(ctx context.Context, id uint64) error
	GetAllCategories(ctx context.Context) ([]models.Category, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (d categoryRepository) CreateCategory(ctx context.Context, category models.Category) (uint64, error) {
	q := `INSERT INTO categories (name, created_at, updated_at) VALUES (?, NOW(), NOW())`
	result, err := d.db.ExecContext(ctx, q, category.Name)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	category.ID = uint64(id)
	return category.ID, nil
}

func (d categoryRepository) GetCategoryByID(ctx context.Context, id uint64) (models.Category, error) {
	var category models.Category
	q := `SELECT id, name, created_at, updated_at FROM categories WHERE id = ?`
	row := d.db.QueryRowContext(ctx, q, id)
	if err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (d categoryRepository) UpdateCategory(ctx context.Context, id uint64, category models.Category) (uint64, error) {
	q := `UPDATE categories SET name = ?, updated_at = NOW() WHERE id = ?`
	result, err := d.db.ExecContext(ctx, q, category.Name, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(rowsAffected), nil
}

func (d categoryRepository) DeleteCategory(ctx context.Context, id uint64) error {
	q := `DELETE FROM categories WHERE id = ?`
	result, err := d.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (d categoryRepository) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	q := `SELECT id, name, created_at, updated_at FROM categories`
	rows, err := d.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
