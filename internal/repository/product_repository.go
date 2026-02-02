package repository

import (
	"context"
	"database/sql"
	"go-crud/internal/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product) (uint64, error)
	GetProductByID(ctx context.Context, id uint64) (models.Product, error)
	UpdateProduct(ctx context.Context, id uint64, product models.Product) (uint64, error)
	DeleteProduct(ctx context.Context, id uint64) error
	GetAllProducts(ctx context.Context) ([]models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (d productRepository) CreateProduct(ctx context.Context, product models.Product) (uint64, error) {
	q := `INSERT INTO products (name, price, category_id, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`
	result, err := d.db.ExecContext(ctx, q, product.Name, product.Price, product.CategoryID)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	product.ID = uint64(id)
	return product.ID, nil
}

func (d productRepository) GetProductByID(ctx context.Context, id uint64) (models.Product, error) {
	var product models.Product
	q := `SELECT id, name, price, category_id, created_at, updated_at FROM products WHERE id = ?`
	row := d.db.QueryRowContext(ctx, q, id)
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (d productRepository) UpdateProduct(ctx context.Context, id uint64, product models.Product) (uint64, error) {
	q := `UPDATE products SET name = ?, price = ?, category_id = ?, updated_at = NOW() WHERE id = ?`
	result, err := d.db.ExecContext(ctx, q, product.Name, product.Price, product.CategoryID, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(rowsAffected), nil
}

func (d productRepository) DeleteProduct(ctx context.Context, id uint64) error {
	q := `DELETE FROM products WHERE id = ?`
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

func (d productRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	q := `SELECT id, name, price, category_id, created_at, updated_at FROM products`
	rows, err := d.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}