package validation

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

type ProductValidation struct {
	Name       string `json:"name"`
	Price      string `json:"price"`
	CategoryID string `json:"category_id"`
}

func ValidateProduct(product ProductValidation) error {
	product.Name = strings.TrimSpace(product.Name)
	if product.Name == "" {
		return errors.New("name is required")
	}

	// Validate and parse price
	product.Price = strings.TrimSpace(product.Price)
	if product.Price == "" {
		return errors.New("price is required")
	}
	price, err := decimal.NewFromString(product.Price)
	if err != nil {
		return errors.New("price must be a valid number")
	}
	if price.LessThanOrEqual(decimal.Zero) {
		return errors.New("price must be greater than 0")
	}

	// Validate and parse category_id
	product.CategoryID = strings.TrimSpace(product.CategoryID)
	if product.CategoryID == "" {
		return errors.New("category_id is required")
	}

	return nil
}
