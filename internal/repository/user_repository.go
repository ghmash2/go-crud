package repository

import (
	"context"
	"database/sql"
	"go-crud/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (uint64, error)
	GetUserByID(ctx context.Context, id uint64) (models.User, error)
	UpdateUser(ctx context.Context, id uint64, user models.User) (uint64, error)
	DeleteUser(ctx context.Context, id uint64) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type dataBase struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &dataBase{db: db}
}

func (d dataBase) CreateUser(ctx context.Context, user models.User) (uint64, error) {
	q := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`
	result, err := d.db.ExecContext(ctx, q, user.Name, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	user.ID = uint64(id)
	return user.ID, nil
}

func (d dataBase) GetUserByID(ctx context.Context, id uint64) (models.User, error) {
	var user models.User
	q := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?`
	row := d.db.QueryRowContext(ctx, q, id)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d dataBase) UpdateUser(ctx context.Context, id uint64, user models.User) (uint64, error) {
	q := `UPDATE users SET name = ?, email = ?, password = ?, updated_at = NOW() WHERE id = ?`
	result, err := d.db.ExecContext(ctx, q, user.Name, user.Email, user.Password, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(rowsAffected), nil
}

func (d dataBase) DeleteUser(ctx context.Context, id uint64) error {
	q := `DELETE FROM users WHERE id = ?`
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

func (d dataBase) GetAllUsers(ctx context.Context) ([]models.User, error) {
	q := `SELECT id, name, email, password, created_at, updated_at FROM users`
	rows, err := d.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}
