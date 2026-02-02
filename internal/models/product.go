package models

import (
	"time"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Price decimal.Decimal `json:"price"`
	CategoryID uint64 `json:"category_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}